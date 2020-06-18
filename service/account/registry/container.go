package registry

import (
    "github.com/jinzhu/gorm"
    "github.com/sarulabs/di/v2"
    "google.golang.org/grpc"
    "google.golang.org/grpc/balancer/roundrobin"

    "github.com/rs/zerolog/log"

    account_entities "github.com/xmlking/grpc-starter-kit/mkit/service/account/entities/v1"
    greeterv1 "github.com/xmlking/grpc-starter-kit/mkit/service/greeter/v1"
    "github.com/xmlking/grpc-starter-kit/service/account/handler"
    "github.com/xmlking/grpc-starter-kit/service/account/repository"
    "github.com/xmlking/grpc-starter-kit/shared/database"
    "github.com/xmlking/grpc-starter-kit/shared/eventing"
    configPB "github.com/xmlking/grpc-starter-kit/shared/proto/config"
)

// Container - provide di Container
type Container struct {
    ctn di.Container
}

// NewContainer - create new Container
func NewContainer(cfg configPB.Configuration) (*Container, error) {
    builder, err := di.NewBuilder()
    if err != nil {
        log.Fatal().Err(err).Msg("")
        return nil, err
    }

    if err := builder.Add([]di.Def{
        {
            Name:  "config",
            Scope: di.App,
            Build: func(ctn di.Container) (interface{}, error) {
                return cfg, nil
            },
        },
        {
            Name:  "user-repository",
            Scope: di.App,
            Build: buildUserRepository,
        },
        {
            Name:  "profile-repository",
            Scope: di.App,
            Build: buildProfileRepository,
        },
        {
            Name:  "email-publisher",
            Scope: di.App,
            Build: func(ctn di.Container) (interface{}, error) {
                return eventing.NewSourceClient(cfg.Services.Emailer.Endpoint), nil
            },
        },
        {
            Name:  "greeter-connection",
            Scope: di.App,
            Build: func(ctn di.Container) (greeterConn interface{}, err error) {
                greeterConn, err = grpc.Dial(cfg.Services.Greeter.Endpoint, grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
                if err != nil {
                    log.Fatal().Msgf("did not connect: %s", err)
                }
                return
            },
            Close: func(obj interface{}) error {
                return obj.(*grpc.ClientConn).Close()
            },
        },
        {
            Name:  "greeter-client",
            Scope: di.App,
            Build: func(ctn di.Container) (interface{}, error) {
                greeterConn := ctn.Get("greeter-connection").(*grpc.ClientConn)
                return greeterv1.NewGreeterServiceClient(greeterConn), nil
            },
        },
        {
            Name:  "user-handler",
            Scope: di.App,
            Build: buildUserHandler,
        },
        {
            Name:  "profile-handler",
            Scope: di.App,
            Build: func(ctn di.Container) (interface{}, error) {
                repo := ctn.Get("profile-repository").(repository.ProfileRepository)

                subLogger := log.With().Str("component", "ProfileHandler").Logger()
                return handler.NewProfileHandler(repo, subLogger), nil
            },
        },
        {
            Name:  "database",
            Scope: di.App,
            Build: func(ctn di.Container) (interface{}, error) {
                return database.GetDatabaseConnection(*cfg.Database)
            },
            Close: func(obj interface{}) error {
                return obj.(*gorm.DB).Close()
            },
        },
    }...); err != nil {
        return nil, err
    }

    return &Container{
        ctn: builder.Build(),
    }, nil
}

// Resolve object
func (c *Container) Resolve(name string) interface{} {
    return c.ctn.Get(name)
}

// Clean Container
func (c *Container) Clean() error {
    return c.ctn.Clean()
}

// Delete Container
func (c *Container) Delete() error {
    return c.ctn.Delete()
}

func buildUserRepository(ctn di.Container) (interface{}, error) {
    db := ctn.Get("database").(*gorm.DB)
    db.AutoMigrate(&account_entities.UserORM{})
    return repository.NewUserRepository(db), nil
}

func buildProfileRepository(ctn di.Container) (interface{}, error) {
    db := ctn.Get("database").(*gorm.DB)
    db.AutoMigrate(&account_entities.ProfileORM{})
    return repository.NewProfileRepository(db), nil
}

func buildUserHandler(ctn di.Container) (interface{}, error) {
    repo := ctn.Get("user-repository").(repository.UserRepository)
    return handler.NewUserHandler(repo, nil, nil), nil // FIXME inject Publisher, and greeter service
}
