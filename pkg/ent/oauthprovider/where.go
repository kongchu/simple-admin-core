// Code generated by ent, DO NOT EDIT.

package oauthprovider

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// ClientID applies equality check predicate on the "client_id" field. It's identical to ClientIDEQ.
func ClientID(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientID), v))
	})
}

// ClientSecret applies equality check predicate on the "client_secret" field. It's identical to ClientSecretEQ.
func ClientSecret(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientSecret), v))
	})
}

// RedirectURL applies equality check predicate on the "redirect_url" field. It's identical to RedirectURLEQ.
func RedirectURL(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRedirectURL), v))
	})
}

// Scopes applies equality check predicate on the "scopes" field. It's identical to ScopesEQ.
func Scopes(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScopes), v))
	})
}

// AuthURL applies equality check predicate on the "auth_url" field. It's identical to AuthURLEQ.
func AuthURL(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthURL), v))
	})
}

// TokenURL applies equality check predicate on the "token_url" field. It's identical to TokenURLEQ.
func TokenURL(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenURL), v))
	})
}

// AuthStyle applies equality check predicate on the "auth_style" field. It's identical to AuthStyleEQ.
func AuthStyle(v uint64) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthStyle), v))
	})
}

// InfoURL applies equality check predicate on the "info_url" field. It's identical to InfoURLEQ.
func InfoURL(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInfoURL), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ClientIDEQ applies the EQ predicate on the "client_id" field.
func ClientIDEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientID), v))
	})
}

// ClientIDNEQ applies the NEQ predicate on the "client_id" field.
func ClientIDNEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClientID), v))
	})
}

// ClientIDIn applies the In predicate on the "client_id" field.
func ClientIDIn(vs ...string) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldClientID), v...))
	})
}

// ClientIDNotIn applies the NotIn predicate on the "client_id" field.
func ClientIDNotIn(vs ...string) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldClientID), v...))
	})
}

// ClientIDGT applies the GT predicate on the "client_id" field.
func ClientIDGT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClientID), v))
	})
}

// ClientIDGTE applies the GTE predicate on the "client_id" field.
func ClientIDGTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClientID), v))
	})
}

// ClientIDLT applies the LT predicate on the "client_id" field.
func ClientIDLT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClientID), v))
	})
}

// ClientIDLTE applies the LTE predicate on the "client_id" field.
func ClientIDLTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClientID), v))
	})
}

// ClientIDContains applies the Contains predicate on the "client_id" field.
func ClientIDContains(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldClientID), v))
	})
}

// ClientIDHasPrefix applies the HasPrefix predicate on the "client_id" field.
func ClientIDHasPrefix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldClientID), v))
	})
}

// ClientIDHasSuffix applies the HasSuffix predicate on the "client_id" field.
func ClientIDHasSuffix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldClientID), v))
	})
}

// ClientIDEqualFold applies the EqualFold predicate on the "client_id" field.
func ClientIDEqualFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldClientID), v))
	})
}

// ClientIDContainsFold applies the ContainsFold predicate on the "client_id" field.
func ClientIDContainsFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldClientID), v))
	})
}

// ClientSecretEQ applies the EQ predicate on the "client_secret" field.
func ClientSecretEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientSecret), v))
	})
}

// ClientSecretNEQ applies the NEQ predicate on the "client_secret" field.
func ClientSecretNEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClientSecret), v))
	})
}

// ClientSecretIn applies the In predicate on the "client_secret" field.
func ClientSecretIn(vs ...string) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldClientSecret), v...))
	})
}

// ClientSecretNotIn applies the NotIn predicate on the "client_secret" field.
func ClientSecretNotIn(vs ...string) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldClientSecret), v...))
	})
}

// ClientSecretGT applies the GT predicate on the "client_secret" field.
func ClientSecretGT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClientSecret), v))
	})
}

// ClientSecretGTE applies the GTE predicate on the "client_secret" field.
func ClientSecretGTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClientSecret), v))
	})
}

// ClientSecretLT applies the LT predicate on the "client_secret" field.
func ClientSecretLT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClientSecret), v))
	})
}

// ClientSecretLTE applies the LTE predicate on the "client_secret" field.
func ClientSecretLTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClientSecret), v))
	})
}

// ClientSecretContains applies the Contains predicate on the "client_secret" field.
func ClientSecretContains(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldClientSecret), v))
	})
}

// ClientSecretHasPrefix applies the HasPrefix predicate on the "client_secret" field.
func ClientSecretHasPrefix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldClientSecret), v))
	})
}

// ClientSecretHasSuffix applies the HasSuffix predicate on the "client_secret" field.
func ClientSecretHasSuffix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldClientSecret), v))
	})
}

// ClientSecretEqualFold applies the EqualFold predicate on the "client_secret" field.
func ClientSecretEqualFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldClientSecret), v))
	})
}

// ClientSecretContainsFold applies the ContainsFold predicate on the "client_secret" field.
func ClientSecretContainsFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldClientSecret), v))
	})
}

// RedirectURLEQ applies the EQ predicate on the "redirect_url" field.
func RedirectURLEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLNEQ applies the NEQ predicate on the "redirect_url" field.
func RedirectURLNEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLIn applies the In predicate on the "redirect_url" field.
func RedirectURLIn(vs ...string) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRedirectURL), v...))
	})
}

// RedirectURLNotIn applies the NotIn predicate on the "redirect_url" field.
func RedirectURLNotIn(vs ...string) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRedirectURL), v...))
	})
}

// RedirectURLGT applies the GT predicate on the "redirect_url" field.
func RedirectURLGT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLGTE applies the GTE predicate on the "redirect_url" field.
func RedirectURLGTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLLT applies the LT predicate on the "redirect_url" field.
func RedirectURLLT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLLTE applies the LTE predicate on the "redirect_url" field.
func RedirectURLLTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLContains applies the Contains predicate on the "redirect_url" field.
func RedirectURLContains(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLHasPrefix applies the HasPrefix predicate on the "redirect_url" field.
func RedirectURLHasPrefix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLHasSuffix applies the HasSuffix predicate on the "redirect_url" field.
func RedirectURLHasSuffix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLEqualFold applies the EqualFold predicate on the "redirect_url" field.
func RedirectURLEqualFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLContainsFold applies the ContainsFold predicate on the "redirect_url" field.
func RedirectURLContainsFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRedirectURL), v))
	})
}

// ScopesEQ applies the EQ predicate on the "scopes" field.
func ScopesEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScopes), v))
	})
}

// ScopesNEQ applies the NEQ predicate on the "scopes" field.
func ScopesNEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldScopes), v))
	})
}

// ScopesIn applies the In predicate on the "scopes" field.
func ScopesIn(vs ...string) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldScopes), v...))
	})
}

// ScopesNotIn applies the NotIn predicate on the "scopes" field.
func ScopesNotIn(vs ...string) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldScopes), v...))
	})
}

// ScopesGT applies the GT predicate on the "scopes" field.
func ScopesGT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldScopes), v))
	})
}

// ScopesGTE applies the GTE predicate on the "scopes" field.
func ScopesGTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldScopes), v))
	})
}

// ScopesLT applies the LT predicate on the "scopes" field.
func ScopesLT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldScopes), v))
	})
}

// ScopesLTE applies the LTE predicate on the "scopes" field.
func ScopesLTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldScopes), v))
	})
}

// ScopesContains applies the Contains predicate on the "scopes" field.
func ScopesContains(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldScopes), v))
	})
}

// ScopesHasPrefix applies the HasPrefix predicate on the "scopes" field.
func ScopesHasPrefix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldScopes), v))
	})
}

// ScopesHasSuffix applies the HasSuffix predicate on the "scopes" field.
func ScopesHasSuffix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldScopes), v))
	})
}

// ScopesEqualFold applies the EqualFold predicate on the "scopes" field.
func ScopesEqualFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldScopes), v))
	})
}

// ScopesContainsFold applies the ContainsFold predicate on the "scopes" field.
func ScopesContainsFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldScopes), v))
	})
}

// AuthURLEQ applies the EQ predicate on the "auth_url" field.
func AuthURLEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthURL), v))
	})
}

// AuthURLNEQ applies the NEQ predicate on the "auth_url" field.
func AuthURLNEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAuthURL), v))
	})
}

// AuthURLIn applies the In predicate on the "auth_url" field.
func AuthURLIn(vs ...string) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAuthURL), v...))
	})
}

// AuthURLNotIn applies the NotIn predicate on the "auth_url" field.
func AuthURLNotIn(vs ...string) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAuthURL), v...))
	})
}

// AuthURLGT applies the GT predicate on the "auth_url" field.
func AuthURLGT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAuthURL), v))
	})
}

// AuthURLGTE applies the GTE predicate on the "auth_url" field.
func AuthURLGTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAuthURL), v))
	})
}

// AuthURLLT applies the LT predicate on the "auth_url" field.
func AuthURLLT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAuthURL), v))
	})
}

// AuthURLLTE applies the LTE predicate on the "auth_url" field.
func AuthURLLTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAuthURL), v))
	})
}

// AuthURLContains applies the Contains predicate on the "auth_url" field.
func AuthURLContains(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAuthURL), v))
	})
}

// AuthURLHasPrefix applies the HasPrefix predicate on the "auth_url" field.
func AuthURLHasPrefix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAuthURL), v))
	})
}

// AuthURLHasSuffix applies the HasSuffix predicate on the "auth_url" field.
func AuthURLHasSuffix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAuthURL), v))
	})
}

// AuthURLEqualFold applies the EqualFold predicate on the "auth_url" field.
func AuthURLEqualFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAuthURL), v))
	})
}

// AuthURLContainsFold applies the ContainsFold predicate on the "auth_url" field.
func AuthURLContainsFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAuthURL), v))
	})
}

// TokenURLEQ applies the EQ predicate on the "token_url" field.
func TokenURLEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenURL), v))
	})
}

// TokenURLNEQ applies the NEQ predicate on the "token_url" field.
func TokenURLNEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTokenURL), v))
	})
}

// TokenURLIn applies the In predicate on the "token_url" field.
func TokenURLIn(vs ...string) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTokenURL), v...))
	})
}

// TokenURLNotIn applies the NotIn predicate on the "token_url" field.
func TokenURLNotIn(vs ...string) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTokenURL), v...))
	})
}

// TokenURLGT applies the GT predicate on the "token_url" field.
func TokenURLGT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTokenURL), v))
	})
}

// TokenURLGTE applies the GTE predicate on the "token_url" field.
func TokenURLGTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTokenURL), v))
	})
}

// TokenURLLT applies the LT predicate on the "token_url" field.
func TokenURLLT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTokenURL), v))
	})
}

// TokenURLLTE applies the LTE predicate on the "token_url" field.
func TokenURLLTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTokenURL), v))
	})
}

// TokenURLContains applies the Contains predicate on the "token_url" field.
func TokenURLContains(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTokenURL), v))
	})
}

// TokenURLHasPrefix applies the HasPrefix predicate on the "token_url" field.
func TokenURLHasPrefix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTokenURL), v))
	})
}

// TokenURLHasSuffix applies the HasSuffix predicate on the "token_url" field.
func TokenURLHasSuffix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTokenURL), v))
	})
}

// TokenURLEqualFold applies the EqualFold predicate on the "token_url" field.
func TokenURLEqualFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTokenURL), v))
	})
}

// TokenURLContainsFold applies the ContainsFold predicate on the "token_url" field.
func TokenURLContainsFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTokenURL), v))
	})
}

// AuthStyleEQ applies the EQ predicate on the "auth_style" field.
func AuthStyleEQ(v uint64) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthStyle), v))
	})
}

// AuthStyleNEQ applies the NEQ predicate on the "auth_style" field.
func AuthStyleNEQ(v uint64) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAuthStyle), v))
	})
}

// AuthStyleIn applies the In predicate on the "auth_style" field.
func AuthStyleIn(vs ...uint64) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAuthStyle), v...))
	})
}

// AuthStyleNotIn applies the NotIn predicate on the "auth_style" field.
func AuthStyleNotIn(vs ...uint64) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAuthStyle), v...))
	})
}

// AuthStyleGT applies the GT predicate on the "auth_style" field.
func AuthStyleGT(v uint64) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAuthStyle), v))
	})
}

// AuthStyleGTE applies the GTE predicate on the "auth_style" field.
func AuthStyleGTE(v uint64) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAuthStyle), v))
	})
}

// AuthStyleLT applies the LT predicate on the "auth_style" field.
func AuthStyleLT(v uint64) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAuthStyle), v))
	})
}

// AuthStyleLTE applies the LTE predicate on the "auth_style" field.
func AuthStyleLTE(v uint64) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAuthStyle), v))
	})
}

// InfoURLEQ applies the EQ predicate on the "info_url" field.
func InfoURLEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInfoURL), v))
	})
}

// InfoURLNEQ applies the NEQ predicate on the "info_url" field.
func InfoURLNEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInfoURL), v))
	})
}

// InfoURLIn applies the In predicate on the "info_url" field.
func InfoURLIn(vs ...string) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldInfoURL), v...))
	})
}

// InfoURLNotIn applies the NotIn predicate on the "info_url" field.
func InfoURLNotIn(vs ...string) predicate.OauthProvider {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldInfoURL), v...))
	})
}

// InfoURLGT applies the GT predicate on the "info_url" field.
func InfoURLGT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInfoURL), v))
	})
}

// InfoURLGTE applies the GTE predicate on the "info_url" field.
func InfoURLGTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInfoURL), v))
	})
}

// InfoURLLT applies the LT predicate on the "info_url" field.
func InfoURLLT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInfoURL), v))
	})
}

// InfoURLLTE applies the LTE predicate on the "info_url" field.
func InfoURLLTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInfoURL), v))
	})
}

// InfoURLContains applies the Contains predicate on the "info_url" field.
func InfoURLContains(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInfoURL), v))
	})
}

// InfoURLHasPrefix applies the HasPrefix predicate on the "info_url" field.
func InfoURLHasPrefix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInfoURL), v))
	})
}

// InfoURLHasSuffix applies the HasSuffix predicate on the "info_url" field.
func InfoURLHasSuffix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInfoURL), v))
	})
}

// InfoURLEqualFold applies the EqualFold predicate on the "info_url" field.
func InfoURLEqualFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInfoURL), v))
	})
}

// InfoURLContainsFold applies the ContainsFold predicate on the "info_url" field.
func InfoURLContainsFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInfoURL), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OauthProvider) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OauthProvider) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OauthProvider) predicate.OauthProvider {
	return predicate.OauthProvider(func(s *sql.Selector) {
		p(s.Not())
	})
}
