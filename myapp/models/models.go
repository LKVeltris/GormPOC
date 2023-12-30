package models

type User struct {
    ID             int    `gorm:"primaryKey"`
    Name           string
    Email          string
    OrganizationID string
    Settings       string
}

type Organization struct {
    ID       int    `gorm:"primaryKey"`
    Name     string
    Settings string
}

type Router struct {
    ID             int    `gorm:"primaryKey"`
    UserID         int
    OrganizationID int
    Config         string
    Status         string
}

type Billing struct {
    ID             int    `gorm:"primaryKey"`
    UserID         int
    OrganizationID int
    BillingInfo    string
    LicenseID      string
}
