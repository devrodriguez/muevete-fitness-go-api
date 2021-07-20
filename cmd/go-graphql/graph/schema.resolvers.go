package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/devrodriguez/muevete-fitness-go-api/cmd/go-graphql/graph/generated"
	"github.com/devrodriguez/muevete-fitness-go-api/cmd/go-graphql/graph/model"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/categories"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/customers"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/routines"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/sessions"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	var category domain.Category
	var newCategory model.Category

	mctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	client, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	category.Name = input.Name

	newCategory.Name = input.Name

	catRepo := dbmongo.NewDbCategoryCrud(client)
	catUc := categories.NewCategoryCrud(catRepo)

	if err := catUc.CreateCategory(ctx, category); err != nil {
		return nil, err
	}

	return &newCategory, nil
}

func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.NewCustomer) (*model.Customer, error) {
	var customer domain.Customer
	var inpCustomer model.Customer

	mctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	client, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	customer.Name = input.Name
	customer.LastName = input.LastName
	customer.Email = input.Email

	inpCustomer.Name = input.Name
	inpCustomer.LastName = input.LastName
	inpCustomer.Email = input.Email

	cusRepo := dbmongo.NewDbCustomerCrud(client)
	cusUc := customers.NewCustomerCrud(cusRepo)

	if err := cusUc.CreateCustomer(ctx, customer); err != nil {
		return nil, err
	}

	return &inpCustomer, nil
}

func (r *mutationResolver) CreateRoutine(ctx context.Context, input model.NewRoutine) (*model.Routine, error) {
	var routine domain.Routine
	var newRoutine model.Routine

	mctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	client, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	routine.Name = input.Name
	routine.Description = input.Description

	newRoutine.Name = input.Name
	newRoutine.Description = input.Description

	rouRepo := dbmongo.NewDbRoutineCrud(client)
	rouUc := routines.NewCrudRoutine(rouRepo)

	if err := rouUc.CreateRoutine(ctx, routine); err != nil {
		return nil, err
	}

	return &newRoutine, nil
}

func (r *mutationResolver) CreateRoutineSchedule(ctx context.Context, input model.NewRoutineSchedule) (*model.RoutineSchedule, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSession(ctx context.Context, input model.NewSession) (*model.Session, error) {
	var session domain.Session
	var newSession model.Session

	mctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	client, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	session.Name = input.Name
	session.Period = input.Period
	session.StartHour = input.StartHour
	session.FinalHour = input.FinalHour

	newSession.Name = input.Name
	newSession.Period = input.Period
	newSession.StartHour = input.StartHour
	newSession.FinalHour = input.FinalHour

	sesRepo := dbmongo.NewDbSessionCrud(client)
	sesUc := sessions.NewCrudSession(sesRepo)

	if err := sesUc.CreateSession(ctx, session); err != nil {
		return nil, err
	}

	return &newSession, nil
}

func (r *mutationResolver) CreateWeekDay(ctx context.Context, input model.NewWeekDay) (*model.WeekDay, error) {
	var weekDay model.WeekDay
	var newWeekDay domain.WeekDay

	mctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	_, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	weekDay.Name = input.Name
	weekDay.NumericDay = input.NumericDay

	newWeekDay.Name = input.Name
	newWeekDay.NumericDay = input.NumericDay

	return nil, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	var qCategories = make([]*model.Category, 0, 10)

	mctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	client, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	catRepo := dbmongo.NewDbCategoryCrud(client)
	catUc := categories.NewCategoryCrud(catRepo)

	data, err := catUc.GetAllCategories(ctx)
	if err != nil {
		return nil, errors.New("error getting categories")
	}

	for _, r := range data {
		catItem := model.Category{
			Name: r.Name,
		}
		qCategories = append(qCategories, &catItem)
	}

	return qCategories, nil
}

func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	var qCustomers = make([]*model.Customer, 0, 10)

	mctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	client, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	cusRepo := dbmongo.NewDbCustomerCrud(client)
	cusUc := customers.NewCustomerCrud(cusRepo)

	cust, err := cusUc.GetAllCustomers(ctx)
	if err != nil {
		return nil, errors.New("error getting customers")
	}

	for _, c := range cust {
		cusItem := model.Customer{
			Name:     c.Name,
			LastName: c.LastName,
			Email:    c.Email,
		}
		qCustomers = append(qCustomers, &cusItem)
	}

	return qCustomers, nil
}

func (r *queryResolver) Routines(ctx context.Context) ([]*model.Routine, error) {
	var qRoutines = make([]*model.Routine, 0, 10)

	mctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	client, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	rouRepo := dbmongo.NewDbRoutineCrud(client)
	rouUc := routines.NewCrudRoutine(rouRepo)

	data, err := rouUc.GetAllRoutines(ctx)
	if err != nil {
		return nil, errors.New("error getting routines")
	}

	for _, r := range data {
		rouItem := model.Routine{
			Name:        r.Name,
			Description: r.Description,
		}
		qRoutines = append(qRoutines, &rouItem)
	}

	return qRoutines, nil
}

func (r *queryResolver) Sessions(ctx context.Context) ([]*model.Session, error) {
	var qSessions = make([]*model.Session, 0, 10)

	mctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	client, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	sesRepo := dbmongo.NewDbSessionCrud(client)
	sesUc := sessions.NewCrudSession(sesRepo)

	data, err := sesUc.GetAllSessions(ctx)

	for _, v := range data {
		sesItem := model.Session{
			Name:      v.Name,
			Period:    v.Period,
			StartHour: v.StartHour,
			FinalHour: v.FinalHour,
		}

		qSessions = append(qSessions, &sesItem)
	}

	return qSessions, nil
}

func (r *queryResolver) WeekDays(ctx context.Context) ([]*model.WeekDay, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) RoutineSchedules(ctx context.Context) ([]*model.RoutineSchedule, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
