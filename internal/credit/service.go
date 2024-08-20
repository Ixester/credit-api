package credit

func GetAllApplicationsService() []CreditApplication {
    return GetAllApplications()
}

func GetApplicationByIDService(id int) *CreditApplication {
    return GetApplicationByID(id)
}

func CreateApplicationService(application CreditApplication) CreditApplication {
    return CreateApplication(application)
}

func ApproveApplicationService(id int, approvedAmount float64) *CreditApplication {
    return ApproveApplication(id, approvedAmount)
}

func RejectApplicationService(id int) *CreditApplication {
    return RejectApplication(id)
}
