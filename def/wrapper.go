package def

import (
	"context"

	cfgDef "github.com/requiemofthesouls/config/def"
	"github.com/requiemofthesouls/container"
	postgres2 "postgres"
)

const (
	DIWrapper      = "postgres.wrapper"
	DIWrapperSqlDB = "postgres.wrapper.sql_db"
)

type (
	Wrapper      = postgres2.Wrapper
	WrapperSqlDB = *postgres2.SqlDB
)

func init() {
	container.Register(func(builder *container.Builder, params map[string]interface{}) error {
		return builder.Add(
			container.Def{
				Name: DIWrapper,
				Build: func(container container.Container) (interface{}, error) {
					var cfg cfgDef.Wrapper
					if err := container.Fill(cfgDef.DIWrapper, &cfg); err != nil {
						return nil, err
					}

					var pgCfg postgres2.Config
					if err := cfg.UnmarshalKey("postgres", &pgCfg); err != nil {
						return nil, err
					}

					return postgres2.New(context.Background(), pgCfg)
				},
				Close: func(obj interface{}) error {
					obj.(postgres2.Wrapper).Close()
					return nil
				},
			},
			container.Def{
				Name: DIWrapperSqlDB,
				Build: func(container container.Container) (interface{}, error) {
					var pgWrapper postgres2.Wrapper
					if err := container.Fill(DIWrapper, &pgWrapper); err != nil {
						return nil, err
					}

					return postgres2.NewSqlDB(pgWrapper)
				},
				Close: func(obj interface{}) error {
					_ = obj.(*postgres2.SqlDB).Close()
					return nil
				},
			})
	})
}
