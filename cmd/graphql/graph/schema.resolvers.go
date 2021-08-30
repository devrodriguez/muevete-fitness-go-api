package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"time"

	"github.com/devrodriguez/muevete-fitness-go-api/cmd/graphql/graph/generated"
	"github.com/devrodriguez/muevete-fitness-go-api/cmd/graphql/graph/model"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/categories"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/customers"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/routines"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/sessions"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/weeklies"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var routineSch domain.RoutineScheduleMod

	mctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	client, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	routineID, _ := primitive.ObjectIDFromHex(input.Routine)
	weekDayID, _ := primitive.ObjectIDFromHex(input.WeekDay)

	routineSch.Routine = routineID
	routineSch.WeekDay = weekDayID

	repo := dbmongo.NewDbRoutineSchedule(client)
	uc := routines.NewRoutineSchedule(repo)

	if err := uc.CreateSchedule(ctx, routineSch); err != nil {
		return nil, err
	}

	return &model.RoutineSchedule{
		Routine: &model.Routine{},
		WeekDay: &model.WeekDay{},
	}, nil
}

func (r *mutationResolver) CreateSessionSchedule(ctx context.Context, input model.NewSessionSchedule) (*model.SessionSchedule, error) {
	var sessionSch domain.SessionScheduleMod

	mctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	client, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	customerID, _ := primitive.ObjectIDFromHex(input.Customer)
	weeklyID, _ := primitive.ObjectIDFromHex(input.Weekly)

	sessionSch.CustomerID = customerID
	sessionSch.WeeklyID = weeklyID

	repo := dbmongo.NewDbSessionSchedule(client)
	uc := sessions.NewSessionSchedule(repo)

	if err := uc.CreateSessionsSchedule(ctx, sessionSch); err != nil {
		return nil, err
	}

	return nil, nil
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

	client, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	newWeekDay.Name = input.Name
	newWeekDay.NumericDay = input.NumericDay

	wdRepo := dbmongo.NewDBWeekDayCrud(client)
	wdUC := weeklies.NewWeekDayCrud(wdRepo)
	wd, err := wdUC.SaveWeekDay(ctx, newWeekDay)
	if err != nil {
		panic(err)
	}

	weekDay.ID = wd.ID.Hex()
	weekDay.Name = input.Name
	weekDay.NumericDay = input.NumericDay

	return &weekDay, nil
}

func (r *mutationResolver) CreateWeekly(ctx context.Context, input model.NewWeekly) (*model.Weekly, error) {
	var weekly model.Weekly
	var newWeekly domain.WeeklyMod

	mctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	client, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	newWeekly.Session, _ = primitive.ObjectIDFromHex(input.Session)
	newWeekly.RoutineSchedule, _ = primitive.ObjectIDFromHex(input.RoutineSchedule)

	repo := dbmongo.NewDbWeeklyCrud(client)
	uc := weeklies.NewWeeklyCrud(repo)

	if err := uc.CreateWeekly(ctx, newWeekly); err != nil {
		return nil, err
	}

	return &weekly, nil
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

	repo := dbmongo.NewDbSessionCrud(client)
	uc := sessions.NewCrudSession(repo)

	data, err := uc.GetAllSessions(ctx)

	if err != nil {
		return nil, err
	}

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
	var qWeekDays = make([]*model.WeekDay, 0, 10)

	mctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	cli, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	// TODO: implements service
	repo := dbmongo.NewDBWeekDayCrud(cli)
	uc := weeklies.NewWeekDayCrud(repo)

	data, err := uc.GetAllDays(ctx)

	if err != nil {
		return nil, err
	}

	for _, v := range data {
		wdItem := model.WeekDay{
			ID:         v.ID.Hex(),
			Name:       v.Name,
			NumericDay: v.NumericDay,
		}

		qWeekDays = append(qWeekDays, &wdItem)
	}

	return qWeekDays, nil
}

func (r *queryResolver) Weeklies(ctx context.Context) ([]*model.Weekly, error) {
	var qWeeklies = make([]*model.Weekly, 0, 10)

	mctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	dbCli, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	repo := dbmongo.NewDbWeeklyCrud(dbCli)
	uc := weeklies.NewWeeklyCrud(repo)

	data, err := uc.GetAllWeeklies(ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range data {
		wkItem := model.Weekly{
			ID: v.ID.Hex(),
			Session: &model.Session{
				ID:        v.Session.ID.Hex(),
				Name:      v.Session.Name,
				StartHour: v.Session.StartHour,
				FinalHour: v.Session.FinalHour,
				Period:    v.Session.Period,
			},
			RoutineSchedule: &model.RoutineSchedule{
				ID: v.RoutineSchedule.ID.Hex(),
				Routine: &model.Routine{
					ID:          v.RoutineSchedule.Routine.ID.Hex(),
					Name:        v.RoutineSchedule.Routine.Name,
					Description: v.RoutineSchedule.Routine.Description,
				},
				WeekDay: &model.WeekDay{
					ID:         v.RoutineSchedule.WeekDay.ID.Hex(),
					Name:       v.RoutineSchedule.WeekDay.Name,
					NumericDay: v.RoutineSchedule.WeekDay.NumericDay,
				},
			},
		}

		qWeeklies = append(qWeeklies, &wkItem)
	}

	return qWeeklies, nil
}

func (r *queryResolver) RoutineSchedules(ctx context.Context) ([]*model.RoutineSchedule, error) {
	var qRoutineSch = make([]*model.RoutineSchedule, 0, 10)

	mctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	client, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	repo := dbmongo.NewDbRoutineSchedule(client)
	uc := routines.NewRoutineSchedule(repo)

	data, err := uc.GetSchedule(ctx)

	for _, v := range data {
		item := model.RoutineSchedule{
			Routine: &model.Routine{
				ID:          v.Routine.ID.Hex(),
				Name:        v.Routine.Name,
				Description: v.Routine.Description,
			},
			WeekDay: &model.WeekDay{
				ID:         v.WeekDay.ID.Hex(),
				Name:       v.WeekDay.Name,
				NumericDay: v.WeekDay.NumericDay,
			},
		}

		qRoutineSch = append(qRoutineSch, &item)
	}

	return qRoutineSch, nil
}

func (r *queryResolver) RoutinesByDay(ctx context.Context) ([]*model.RoutineCategory, error) {
	var rcs = make([]*model.RoutineCategory, 0, 10)

	mctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	client, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	repo := dbmongo.NewDbRoutineCrud(client)
	uc := routines.NewCrudRoutine(repo)

	data, err := uc.GetRoutineByDay(ctx, "")
	if err != nil {
		return nil, err
	}

	for _, v := range data {
		item := model.RoutineCategory{
			Category: &model.Category{
				ID:   v.Category.ID.Hex(),
				Name: v.Category.Name,
			},
		}

		for _, v := range v.Routines {
			itemRoutine := &model.Routine{
				ID:          v.ID.Hex(),
				Name:        v.Name,
				Description: v.Description,
			}

			item.Routines = append(item.Routines, itemRoutine)
		}

		rcs = append(rcs, &item)
	}

	return rcs, nil
}

func (r *queryResolver) SessionSchedules(ctx context.Context) ([]*model.SessionSchedule, error) {
	var qSessionSch = make([]*model.SessionSchedule, 0, 10)

	mctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	client, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	repo := dbmongo.NewDbSessionSchedule(client)
	uc := sessions.NewSessionSchedule(repo)

	data, err := uc.GetSchedule(ctx)

	if err != nil {
		return nil, err
	}

	for _, v := range data {
		item := model.SessionSchedule{
			Customer: &model.Customer{
				ID:       v.Customer.ID.Hex(),
				Name:     v.Customer.Name,
				LastName: v.Customer.LastName,
				Email:    v.Customer.Email,
			},
			Weekly: &model.Weekly{
				ID: v.Weekly.ID.Hex(),
				Session: &model.Session{
					ID:        v.Weekly.Session.ID.Hex(),
					Name:      v.Session.Name,
					StartHour: v.Session.StartHour,
					FinalHour: v.Session.FinalHour,
					Period:    v.Session.Period,
				},
				RoutineSchedule: &model.RoutineSchedule{
					ID: v.RoutineSchedule.ID.Hex(),
					Routine: &model.Routine{
						ID:          v.RoutineSchedule.Routine.ID.Hex(),
						Name:        v.RoutineSchedule.Routine.Name,
						Description: v.RoutineSchedule.Routine.Description,
					},
					WeekDay: &model.WeekDay{
						ID:         v.RoutineSchedule.WeekDay.ID.Hex(),
						Name:       v.RoutineSchedule.WeekDay.Name,
						NumericDay: v.RoutineSchedule.WeekDay.NumericDay,
					},
				},
			},
		}

		qSessionSch = append(qSessionSch, &item)
	}

	return qSessionSch, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
