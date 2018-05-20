# Entity Framework Core Console Example

Create Initial Migrations (they will be generated in a new "Migrations" folder):

```
dotnet ef migrations add InitialCreate
```

Apply Migrations:

```
dotnet ef database update
```

Update Db when changing models:

```
dotnet ef migrations add
dotnet ef database update
```

Run

```
dotnet run
```

Additional Information

 - [Testing with SQLite in-memory](https://docs.microsoft.com/en-us/ef/core/miscellaneous/testing/sqlite)
 - [ASP.net Core, getting with a new database](https://docs.microsoft.com/en-us/ef/core/get-started/aspnetcore/new-db)
 - [Console app (this) and migrations](https://www.learnentityframeworkcore.com/walkthroughs/console-application)
 - [Npgsql provider for postgres](http://www.npgsql.org/efcore/index.html)

[&#x2190; Back to Project List](../README.md)