package credit

import "time"

var applications []CreditApplication
var nextID int = 1

func GetAllApplications() []CreditApplication {
    return applications
}

func GetApplicationByID(id int) *CreditApplication {
    for _, application := range applications {
        if application.ID == id {
            return &application
        }
    }
    return nil
}

func CreateApplication(application CreditApplication) CreditApplication {
    application.ID = nextID
    application.Status = "Pending"
    application.ApplicationDate = time.Now()
    nextID++
    applications = append(applications, application)
    return application
}

func ApproveApplication(id int, approvedAmount float64) *CreditApplication {
    for i, application := range applications {
        if application.ID == id {
            application.Status = "Approved"
            application.ApprovedAmount = &approvedAmount
            approvalDate := time.Now()
            application.ApprovalDate = &approvalDate
            applications[i] = application
            return &applications[i]
        }
    }
    return nil
}

func RejectApplication(id int) *CreditApplication {
    for i, application := range applications {
        if application.ID == id {
            application.Status = "Rejected"
            applications[i] = application
            return &applications[i]
        }
    }
    return nil
}
