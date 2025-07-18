package postgres

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	cr "github.com/VoroniakPavlo/call_audit/api/protos/storage"
	dberr "github.com/VoroniakPavlo/call_audit/internal/errors"
	"github.com/VoroniakPavlo/call_audit/internal/store/postgres/scanner"
	"github.com/VoroniakPavlo/call_audit/internal/store/util"
	options "github.com/VoroniakPavlo/call_audit/model/options"
)

type QuestionnaireRuleScan func(rule *cr.CallQuestionnaireRule) any

const (
	cqrLeft                           = "cqr"
	closeQuestionnaireRuleDefaultSort = "name"
)

// CallQuestionnaireRule provides methods to interact with call questionnaire rules in the database.
type CallQuestionnaireRuleStore struct {
	storage *Store
}

// Create implements store.CallQuestionnaireRuleStore.
func (c *CallQuestionnaireRuleStore) Create(ctx context.Context, rule *cr.CallQuestionnaireRule) (*cr.CallQuestionnaireRule, error) {
	panic("unimplemented")
}

// Delete implements store.CallQuestionnaireRuleStore.
func (c *CallQuestionnaireRuleStore) Delete(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// Get implements store.CallQuestionnaireRuleStore.
func (c *CallQuestionnaireRuleStore) Get(ctx context.Context, id int64) (*cr.CallQuestionnaireRule, error) {
	panic("unimplemented")
}

// List implements store.CallQuestionnaireRuleStore.
func (c *CallQuestionnaireRuleStore) List(ctx context.Context) (*cr.CallQuestionnaireRuleList, error) {
	d, dbErr := c.storage.Database()
	if dbErr != nil {
		return nil, dberr.NewDBInternalError("postgres.call_questionnaire_rule.list.database_connection_error", dbErr)
	}

	selectBuilder, plan, err := c.buildListCallQuestionnaireRuleQuery(ctx, closeQuestionnaireRuleId)
	if err != nil {
		return nil, dberr.NewDBInternalError("postgres.call_questionnaire_rule.list.build_query_error", err)
	}

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, dberr.NewDBInternalError("postgres.close_reason.list.query_build_error", err)
	}
	query = util.CompactSQL(query)

	rows, err := d.Query(ctx, query, args...)
	if err != nil {
		return nil, dberr.NewDBInternalError("postgres.close_reason.list.execution_error", err)
	}
	defer rows.Close()

	var reasons []*cr.CallQuestionnaireRule
	lCount := 0
	next := false
	fetchAll := ctx.GetSize() == -1

	for rows.Next() {
		if !fetchAll && lCount >= int(ctx.GetSize()) {
			next = true
			break
		}

		reason := &cr.CallQuestionnaireRule{}
		scanArgs := convertToCloseReasonScanArgs(plan, reason)

		if err := rows.Scan(scanArgs...); err != nil {
			return nil, dberr.NewDBInternalError("postgres.close_reason.list.row_scan_error", err)
		}

		reasons = append(reasons, reason)
		lCount++
	}

	return &cr.CallQuestionnaireRuleList{
		Page:  int32(ctx.GetPage()),
		Next:  next,
		Items: reasons,
	}, nil
}

// Update implements store.CallQuestionnaireRuleStore.
func (c *CallQuestionnaireRuleStore) Update(ctx context.Context, rule *cr.CallQuestionnaireRule) (*cr.CallQuestionnaireRule, error) {
	panic("unimplemented")
}

// NewCallQuestionnaireRuleStore creates a new CallQuestionnaireRuleStore.
func NewCallQuestionnaireRuleStore(storage *Store) *CallQuestionnaireRuleStore {
	return &CallQuestionnaireRuleStore{storage: storage}
}

func (s *CallQuestionnaireRuleStore) buildListCallQuestionnaireRuleQuery(
	rpc options.SearchOptions,
	callQuestionnaireRuleId int64,
) (sq.SelectBuilder, []QuestionnaireRuleScan, error) {
	queryBuilder := sq.Select().
		From("call_audit.call_questionnaire_rule AS cqr").
		Where(sq.Eq{"cqr.domain": rpc.GetAuthOpts().GetDomainId()}).
		PlaceholderFormat(sq.Dollar)

	// Add ID filter if provided
	if len(rpc.GetIDs()) > 0 {
		queryBuilder = queryBuilder.Where(sq.Eq{"cqr.id": rpc.GetIDs()})
	}

	// -------- Apply sorting ----------
	queryBuilder = util.ApplyDefaultSorting(rpc, queryBuilder, closeQuestionnaireRuleDefaultSort)

	// ---------Apply paging based on Search Opts ( page ; size ) -----------------
	queryBuilder = util.ApplyPaging(rpc.GetPage(), rpc.GetSize(), queryBuilder)

	// Add select columns and scan plan for requested fields
	queryBuilder, plan, err := buildQuestionnaireRuleSelectColumnsAndPlan(queryBuilder, rpc.GetFields())
	if err != nil {
		return sq.SelectBuilder{}, nil, dberr.NewDBInternalError("postgres.questionnaire_rule.search.query_build_error", err)
	}

	return queryBuilder, plan, nil
}

func buildQuestionnaireRuleSelectColumnsAndPlan(
	base sq.SelectBuilder,
	fields []string,
) (sq.SelectBuilder, []QuestionnaireRuleScan, error) {
	var plan []QuestionnaireRuleScan
	for _, field := range fields {
		switch field {
		case "id":
			base = base.Column(util.Ident(cqrLeft, "id"))
			plan = append(plan, func(rule *cr.CallQuestionnaireRule) any {
				return &rule.Id
			})
		case "name":
			base = base.Column(util.Ident(cqrLeft, "name"))
			plan = append(plan, func(rule *cr.CallQuestionnaireRule) any {
				return &rule.Name
			})
		case "description":
			base = base.Column(util.Ident(cqrLeft, "description"))
			plan = append(plan, func(rule *cr.CallQuestionnaireRule) any {
				return &rule.Description
			})

		case "created_at":
			base = base.Column(util.Ident(cqrLeft, "created_at"))
			plan = append(plan, func(rule *cr.CallQuestionnaireRule) any {
				return &rule.CreatedAt
			})
		case "updated_at":
			base = base.Column(util.Ident(cqrLeft, "updated_at"))
			plan = append(plan, func(rule *cr.CallQuestionnaireRule) any {
				return &rule.UpdatedAt
			})
		case "updated_by":
			base = base.Column(fmt.Sprintf(
				"(SELECT ROW(id, COALESCE(name, username))::text FROM directory.wbt_user WHERE id = %s.updated_by) updated_by",
				cqrLeft))
			plan = append(plan, func(rule *cr.CallQuestionnaireRule) any {
				return scanner.ScanRowLookup(&rule.UpdatedBy)
			})
		case "created_by":
			base = base.Column(fmt.Sprintf(
				"(SELECT ROW(id, COALESCE(name, username))::text FROM directory.wbt_user WHERE id = %s.created_by) created_by",
				cqrLeft))
			plan = append(plan, func(rule *cr.CallQuestionnaireRule) any {
				return scanner.ScanRowLookup(&rule.CreatedBy)
			})
		case "enabled":
			base = base.Column(util.Ident(cqrLeft, "enabled"))
			plan = append(plan, func(rule *cr.CallQuestionnaireRule) any {
				return &rule.Enabled
			})
		case "domain_id":
			base = base.Column(util.Ident(cqrLeft, "domain_id"))
			plan = append(plan, func(rule *cr.CallQuestionnaireRule) any {
				return &rule.DomainId
			})
		case "last_stored_at":
			base = base.Column(util.Ident(cqrLeft, "last_stored_at"))
			plan = append(plan, func(rule *cr.CallQuestionnaireRule) any {
				return scanner.ScanTimestamp(&rule.LastStoredAt)
			})
		default:
			return base, nil, dberr.NewDBInternalError("postgres.close_reason.unknown_field", fmt.Errorf("unknown field: %s", field))
		}
	}
	return base, plan, nil
}
