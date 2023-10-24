package config

import "idstar-idp/rest-api/docs"

func InitSwagger() {
	docs.SwaggerInfo.Title = getConfigValue("swagger.title")
	docs.SwaggerInfo.Description = getConfigValue("swagger.description")
	docs.SwaggerInfo.Version = getConfigValue("swagger.version")
	docs.SwaggerInfo.Host = getConfigValue("swagger.host")
	docs.SwaggerInfo.BasePath = getConfigValue("swagger.basePath")
	docs.SwaggerInfo.Schemes = []string{"http"}
}
