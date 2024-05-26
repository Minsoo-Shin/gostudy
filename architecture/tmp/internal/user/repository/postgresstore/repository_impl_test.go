package postgresstore

//func TestUserPostgresStore_FindByID(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	gormDB, err := gorm.Open(mysql.New(mysql.Config{
//		Conn:                      db,
//		DriverName:                "mysql",
//		SkipInitializeWithVersion: true,
//	}), &gorm.Config{})
//	if err != nil {
//		log.Fatalf("%v", err)
//	}
//
//	repo := New(gormDB)
//	newUser := domain.User{}
//	err = faker.FakeData(&newUser)
//	if err != nil {
//		fmt.Println(err)
//	}
//	newUser.ID = 1
//
//	mock.ExpectBegin()
//	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`username`,`email`,`password`,`name`,`birthdate`,`id`) VALUES (?,?,?,?,?,?,?,?,?)")).
//		WithArgs(newUser.CreatedAt, newUser.UpdatedAt, newUser.DeletedAt, newUser.Username, newUser.Email, newUser.Password, newUser.Name, newUser.Birthdate, newUser.ID).
//		WillReturnResult(sqlmock.NewResult(1, 1))
//	mock.ExpectCommit()
//	err = repo.Create(context.Background(), newUser)
//
//	require.NoError(t, err)
//	got, err := repo.FindByID(context.Background(), 1)
//	require.NoError(t, err)
//	require.Equalf(t, newUser, got, "want: %v, got: %v", newUser, got)
//}
