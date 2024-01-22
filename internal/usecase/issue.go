package usecase

import (
	"fmt"

	"github.com/google/uuid"
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

type EquipmentWithQuantity struct {
	EquipmentID string `json:"equipmentId"`
	Quantity    int    `json:"quantity"`
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

func equipmentWithQuantityToModelLoanEntry(src EquipmentWithQuantity) model.LoanEntry {
	return *model.NewLoanEntry(int32(src.Quantity), uuid.MustParse(src.EquipmentID), uuid.UUID{})
}

func (iu *IssueUseCase) CreateIssue(address, name, note string, equipments []EquipmentWithQuantity) (uuid.UUID, error) {
	ir := iu.r.GetIssueRepository() //nolint:varnamelen

	loanEntries := lop.Map(equipments, func(data EquipmentWithQuantity, _ int) *model.LoanEntry {
		loanEntry := equipmentWithQuantityToModelLoanEntry(data)

		return &loanEntry
	})

	modelIssue := model.NewIssue(address, name, string(model.StatusSurvey), note, loanEntries)

	err := iu.r.Atomic(func(br repository.BaseRepository) error {
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
