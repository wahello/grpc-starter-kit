// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
	"github.com/xmlking/grpc-starter-kit/ent/profile"
	"github.com/xmlking/grpc-starter-kit/ent/user"
)

// ProfileCreate is the builder for creating a Profile entity.
type ProfileCreate struct {
	config
	mutation *ProfileMutation
	hooks    []Hook
}

// SetCreateTime sets the create_time field.
func (pc *ProfileCreate) SetCreateTime(t time.Time) *ProfileCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (pc *ProfileCreate) SetNillableCreateTime(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the update_time field.
func (pc *ProfileCreate) SetUpdateTime(t time.Time) *ProfileCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (pc *ProfileCreate) SetNillableUpdateTime(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetDeleteTime sets the delete_time field.
func (pc *ProfileCreate) SetDeleteTime(t time.Time) *ProfileCreate {
	pc.mutation.SetDeleteTime(t)
	return pc
}

// SetNillableDeleteTime sets the delete_time field if the given value is not nil.
func (pc *ProfileCreate) SetNillableDeleteTime(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetDeleteTime(*t)
	}
	return pc
}

// SetAge sets the age field.
func (pc *ProfileCreate) SetAge(i int) *ProfileCreate {
	pc.mutation.SetAge(i)
	return pc
}

// SetTz sets the tz field.
func (pc *ProfileCreate) SetTz(s string) *ProfileCreate {
	pc.mutation.SetTz(s)
	return pc
}

// SetAvatar sets the avatar field.
func (pc *ProfileCreate) SetAvatar(u *url.URL) *ProfileCreate {
	pc.mutation.SetAvatar(u)
	return pc
}

// SetBirthday sets the birthday field.
func (pc *ProfileCreate) SetBirthday(t time.Time) *ProfileCreate {
	pc.mutation.SetBirthday(t)
	return pc
}

// SetNillableBirthday sets the birthday field if the given value is not nil.
func (pc *ProfileCreate) SetNillableBirthday(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetBirthday(*t)
	}
	return pc
}

// SetGender sets the gender field.
func (pc *ProfileCreate) SetGender(pr profile.Gender) *ProfileCreate {
	pc.mutation.SetGender(pr)
	return pc
}

// SetNillableGender sets the gender field if the given value is not nil.
func (pc *ProfileCreate) SetNillableGender(pr *profile.Gender) *ProfileCreate {
	if pr != nil {
		pc.SetGender(*pr)
	}
	return pc
}

// SetPreferredTheme sets the preferred_theme field.
func (pc *ProfileCreate) SetPreferredTheme(s string) *ProfileCreate {
	pc.mutation.SetPreferredTheme(s)
	return pc
}

// SetNillablePreferredTheme sets the preferred_theme field if the given value is not nil.
func (pc *ProfileCreate) SetNillablePreferredTheme(s *string) *ProfileCreate {
	if s != nil {
		pc.SetPreferredTheme(*s)
	}
	return pc
}

// SetID sets the id field.
func (pc *ProfileCreate) SetID(u uuid.UUID) *ProfileCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetUserID sets the user edge to User by id.
func (pc *ProfileCreate) SetUserID(id uuid.UUID) *ProfileCreate {
	pc.mutation.SetUserID(id)
	return pc
}

// SetUser sets the user edge to User.
func (pc *ProfileCreate) SetUser(u *User) *ProfileCreate {
	return pc.SetUserID(u.ID)
}

// Mutation returns the ProfileMutation object of the builder.
func (pc *ProfileCreate) Mutation() *ProfileMutation {
	return pc.mutation
}

// Save creates the Profile in the database.
func (pc *ProfileCreate) Save(ctx context.Context) (*Profile, error) {
	if err := pc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Profile
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProfileCreate) SaveX(ctx context.Context) *Profile {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *ProfileCreate) preSave() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		v := profile.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		v := profile.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
	if _, ok := pc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New("ent: missing required field \"age\"")}
	}
	if v, ok := pc.mutation.Age(); ok {
		if err := profile.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Tz(); !ok {
		return &ValidationError{Name: "tz", err: errors.New("ent: missing required field \"tz\"")}
	}
	if v, ok := pc.mutation.Gender(); ok {
		if err := profile.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf("ent: validator failed for field \"gender\": %w", err)}
		}
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := profile.DefaultID()
		pc.mutation.SetID(v)
	}
	if _, ok := pc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New("ent: missing required edge \"user\"")}
	}
	return nil
}

func (pc *ProfileCreate) sqlSave(ctx context.Context) (*Profile, error) {
	pr, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}

func (pc *ProfileCreate) createSpec() (*Profile, *sqlgraph.CreateSpec) {
	var (
		pr    = &Profile{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: profile.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: profile.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		pr.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldCreateTime,
		})
		pr.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldUpdateTime,
		})
		pr.UpdateTime = value
	}
	if value, ok := pc.mutation.DeleteTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldDeleteTime,
		})
		pr.DeleteTime = &value
	}
	if value, ok := pc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profile.FieldAge,
		})
		pr.Age = value
	}
	if value, ok := pc.mutation.Tz(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldTz,
		})
		pr.Tz = value
	}
	if value, ok := pc.mutation.Avatar(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: profile.FieldAvatar,
		})
		pr.Avatar = value
	}
	if value, ok := pc.mutation.Birthday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldBirthday,
		})
		pr.Birthday = value
	}
	if value, ok := pc.mutation.Gender(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: profile.FieldGender,
		})
		pr.Gender = value
	}
	if value, ok := pc.mutation.PreferredTheme(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldPreferredTheme,
		})
		pr.PreferredTheme = value
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profile.UserTable,
			Columns: []string{profile.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pr, _spec
}

// ProfileCreateBulk is the builder for creating a bulk of Profile entities.
type ProfileCreateBulk struct {
	config
	builders []*ProfileCreate
}

// Save creates the Profile entities in the database.
func (pcb *ProfileCreateBulk) Save(ctx context.Context) ([]*Profile, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Profile, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*ProfileMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (pcb *ProfileCreateBulk) SaveX(ctx context.Context) []*Profile {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
