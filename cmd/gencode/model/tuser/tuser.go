
package tuser

type TUser struct{
    ID int 
    Name string 
    UserName string 
    Password string 
    CreatedAt time.Time 
    UpdatedAt time.Time 
    DeletedAt time.Time 
    RememberToken string 
    EmailVerifiedAt time.Time 
}
