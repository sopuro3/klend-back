package usecase

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"

	"github.com/sopuro3/klend-back/internal/model"
	"github.com/sopuro3/klend-back/internal/repository"
)

type (
	IssueID   string // uuidv7
	DisplayID string // 数字4桁
)

type Issue struct {
	Address   string            `json:"address"` // 128文字
	Name      string            `json:"name"`    // 128文字
	IssueID   IssueID           `json:"issueId"`
	DisplayID DisplayID         `json:"displayId"`
	Status    model.IssueStatus `json:"status"`
	Note      string            `json:"note"` // 256文字
}

type EquipmentWithPlannedQuantity struct {
	EquipmentID string `json:"equipmentId"`
	Quantity    int    `json:"plannedQuantity"`
}

type IssueUseCase struct {
	r repository.BaseRepository
}

func NewIssueUseCase(r repository.BaseRepository) *IssueUseCase {
	return &IssueUseCase{r}
}

func modelIssueToIssue(modelIssue *model.Issue) Issue {
	return Issue{
		Address:   modelIssue.Address,
		Name:      modelIssue.Name,
		IssueID:   IssueID(modelIssue.ID.String()),
		DisplayID: DisplayID(fmt.Sprintf("%04d", *modelIssue.DisplayID.ID)),
		Status:    model.IssueStatus(modelIssue.Status),
		Note:      modelIssue.Note,
	}
}

func (iu *IssueUseCase) GetIssueAll() ([]Issue, error) {
	ir := iu.r.GetIssueRepository()

	modelIssues, err := ir.FindAll()
	if err != nil {
		return nil, err
	}

	issues := lop.Map(modelIssues, func(e *model.Issue, _ int) Issue {
		return modelIssueToIssue(e)
	})

	return issues, nil
}

func (iu *IssueUseCase) GetIssue(id uuid.UUID) (Issue, error) {
	ir := iu.r.GetIssueRepository()

	modelIssues, err := ir.Find(id)
	if err != nil {
		return Issue{}, err
	}

	issue := modelIssueToIssue(modelIssues)

	return issue, nil
}

func equipmentWithQuantityToModelLoanEntry(src EquipmentWithPlannedQuantity) model.LoanEntry {
	return *model.NewLoanEntry(int32(src.Quantity), uuid.MustParse(src.EquipmentID), uuid.UUID{})
}

func (iu *IssueUseCase) CreateIssue(
	address, name, note string, equipments []EquipmentWithPlannedQuantity,
) (uuid.UUID, error) {
	loanEntries := lop.Map(equipments, func(data EquipmentWithPlannedQuantity, _ int) *model.LoanEntry {
		loanEntry := equipmentWithQuantityToModelLoanEntry(data)

		return &loanEntry
	})

	modelIssue := model.NewIssue(address, name, string(model.StatusSurvey), note, loanEntries)

	err := iu.r.Atomic(func(br repository.BaseRepository) error {
		ir := br.GetIssueRepository() //nolint:varnamelen
		if err := ir.Create(modelIssue); err != nil {
			return err
		}

		dir := br.GetDisplayIDRepository()

		_, err := dir.Create(modelIssue.ID)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	return modelIssue.ID, nil
}

func (iu *IssueUseCase) UpdateIssue(issueID uuid.UUID, address, name, note *string, equipments []EquipmentWithPlannedQuantity) error {
	var err error

	oldIssue, err := iu.r.GetIssueRepository().Find(issueID)
	if err != nil {
		return ErrRecodeNotFound
	}

	if oldIssue.Status != string(model.StatusSurvey) && oldIssue.Status != string(model.StatusEquipmentCheck) {
		return ErrInvalidStatus
	}

	issue := model.Issue{
		Model:  model.Model{ID: oldIssue.ID},
		Status: string(model.StatusEquipmentCheck),
	}

	if address != nil {
		issue.Address = *address
	}

	if name != nil {
		issue.Name = *name
	}

	if note != nil {
		issue.Note = *note
	}

	oldLoanEntries, err := iu.r.GetLoanEntryRepository().FindByIssueID(issueID)
	if err != nil {
		return ErrRecodeNotFound
	}

	slog.Info("issue 2", "oldLoanEntries", oldLoanEntries)

	loanEntries := make([]*model.LoanEntry, 0, len(oldLoanEntries))

	for _, v := range equipments {
		eqID, err := uuid.Parse(v.EquipmentID)
		if err != nil {
			// TODO 不正なIDの考慮
			continue
		}

		loanEntry, ok := lo.Find(oldLoanEntries, func(item *model.LoanEntry) bool {
			return item.EquipmentID == eqID
		})
		if !ok {
			continue
		}

		loanEntry.Quantity = int32(v.Quantity)
		loanEntries = append(loanEntries, loanEntry)
	}
	slog.Info("issue 5", "loanEntries", loanEntries)

	err = iu.r.Atomic(func(br repository.BaseRepository) error {
		err := br.GetIssueRepository().Update(&issue)
		if err != nil {
			return err
		}

		lr := br.GetLoanEntryRepository()
		for _, v := range loanEntries {
			err := lr.Update(v)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
