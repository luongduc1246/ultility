package gormdb

// func TestGormClause(t *testing.T) {
// 	var config = &gorm.Config{
// 		// Logger: logger.Default.LogMode(logger.Silent),
// 		SkipDefaultTransaction: true,
// 		NamingStrategy: schema.NamingStrategy{
// 			TablePrefix: "account.alib_", // table name prefix, table for `User` would be `goqh_users`

// 		},
// 	}
// 	txtUri, bol := os.LookupEnv("DATABASE_URL")
// 	if !bol {
// 		txtUri = "postgres://mrpink:luongduc7246@localhost:5432/alibaba"
// 	}
// 	instance, _ = gorm.Open(postgres.Open(txtUri), config)
// 	users := []User{}
// 	tx := instance.Session(&gorm.Session{})
// 	tx.Statement.Parse(&User{})
// 	instance.Debug().Clauses(clause.Select{Columns: []clause.Column{clause.PrimaryColumn, {Name: "uuid"}}}, clause.From{
// 		Joins: []clause.Join{{}},
// 	}).Find(&users)
// 	fmt.Printf("%+v \n", tx.Statement.Schema.Table)
// 	tx.Statement.Parse(&User{})
// 	fmt.Printf("%+v \n", tx.Statement.Schema.Relationships.Relations["Roles"].FieldSchema)
// 	// fmt.Printf("%+v \n", tx.Statement.Schema.Relationships.Relations["Users"].References[1].PrimaryKey.Schema.Table)
// }

// func TestParseFieldToFieldSelect(t *testing.T) {
// 	var config = &gorm.Config{
// 		// Logger: logger.Default.LogMode(logger.Silent),
// 		SkipDefaultTransaction: true,
// 		NamingStrategy: schema.NamingStrategy{
// 			TablePrefix: "account.alib_", // table name prefix, table for `User` would be `goqh_users`

// 		},
// 	}
// 	txtUri, bol := os.LookupEnv("DATABASE_URL")
// 	if !bol {
// 		txtUri = "postgres://mrpink:luongduc7246@localhost:5432/alibaba"
// 	}
// 	instance, _ = gorm.Open(postgres.Open(txtUri), config)
// 	// users := []User{}
// 	tx := instance.Session(&gorm.Session{})
// 	tx.Statement.Parse(&User{})
// 	a := "uuid,phone,roles[uuid,name]"
// 	f := reqparams.NewField()
// 	err := f.Parse(a)
// 	fmt.Println(err)
// 	s := newFieldSelect()
// 	s.parse(tx.Statement.Schema, f)
// 	fmt.Println(s)
// }

// func TestParseFilterToFilterWhere(t *testing.T) {
// 	var config = &gorm.Config{
// 		// Logger: logger.Default.LogMode(logger.Silent),
// 		SkipDefaultTransaction: true,
// 		NamingStrategy: schema.NamingStrategy{
// 			TablePrefix: "account.alib_", // table name prefix, table for `User` would be `goqh_users`

// 		},
// 	}
// 	txtUri, bol := os.LookupEnv("DATABASE_URL")
// 	if !bol {
// 		txtUri = "postgres://mrpink:luongduc7246@localhost:5432/alibaba"
// 	}
// 	instance, _ = gorm.Open(postgres.Open(txtUri), config)
// 	// users := []User{}
// 	tx := instance.Session(&gorm.Session{})
// 	tx.Statement.Parse(&User{})
// 	s := `eq[phone]=%25name%25,not[eq[name]=haha,taka[eq[name]=Tetstata,or[and[eq[name]=alfjaakaj]]]],roles[eq[name]=admin,or[eq[id]=3]]`
// 	f := reqparams.NewFilter()
// 	err := f.Parse(s)
// 	fmt.Println(err)
// 	fmt.Println(f)
// 	join := map[string][]clause.Join{}
// 	// fw := NewFilterWhere()
// 	exps := parseFilterToFilterWhere(tx.Statement.Schema, f, join)
// 	fmt.Println(exps)
// 	fmt.Printf("%+v \n", join)
// }
// func TestParseSortToSortBy(t *testing.T) {
// 	var config = &gorm.Config{
// 		// Logger: logger.Default.LogMode(logger.Silent),
// 		SkipDefaultTransaction: true,
// 		NamingStrategy: schema.NamingStrategy{
// 			TablePrefix: "account.alib_", // table name prefix, table for `User` would be `goqh_users`

// 		},
// 	}
// 	txtUri, bol := os.LookupEnv("DATABASE_URL")
// 	if !bol {
// 		txtUri = "postgres://mrpink:luongduc7246@localhost:5432/alibaba"
// 	}
// 	instance, _ = gorm.Open(postgres.Open(txtUri), config)
// 	// users := []User{}
// 	tx := instance.Session(&gorm.Session{})
// 	tx.Statement.Parse(&User{})
// 	s := `asc[phone],desc[id],roles[asc[id]]`
// 	sort := reqparams.NewSort()
// 	err := sort.Parse(s)
// 	fmt.Println(err)
// 	fmt.Println(sort)
// 	join := map[string][]clause.OrderByColumn{}
// 	// fw := NewFilterWhere()
// 	exps := parseSortToSortBy(tx.Statement.Schema, sort, join)
// 	fmt.Println(exps)
// 	fmt.Printf("%+v \n", join)
// }

// func TestParseClauseSearch(t *testing.T) {
// 	var config = &gorm.Config{
// 		// Logger: logger.Default.LogMode(logger.Silent),
// 		SkipDefaultTransaction: true,
// 		NamingStrategy: schema.NamingStrategy{
// 			TablePrefix: "account.alib_", // table name prefix, table for `User` would be `goqh_users`

// 		},
// 	}
// 	txtUri, bol := os.LookupEnv("DATABASE_URL")
// 	if !bol {
// 		txtUri = "postgres://mrpink:luongduc7246@localhost:5432/alibaba"
// 	}
// 	instance, _ = gorm.Open(postgres.Open(txtUri), config)
// 	users := []User{}
// 	tx := instance.Session(&gorm.Session{})
// 	tx.Statement.Parse(&User{})
// 	param := reqparams.Params{
// 		Filter: []string{`like[phone]=%97%`},
// 		Sort:   []string{`asc[phone],desc[id],roles[asc[id]]`},
// 	}
// 	s := reqparams.NewSearch()
// 	err := s.Parse(param)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	cs := NewClauseSearch()
// 	cs.Parse(tx.Statement.Schema, s)
// 	exps := cs.Build()
// 	instance.Debug().Preload("Roles").Clauses(exps...).Find(&users)
// 	fmt.Println(users)
// }
