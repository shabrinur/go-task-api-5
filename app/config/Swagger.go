package config

import "idstar-idp/rest-api/docs"

func InitSwagger() {
	docs.SwaggerInfo.Title = GetConfigValue("swagger.title")
	docs.SwaggerInfo.Description = GetConfigValue("swagger.description")
	docs.SwaggerInfo.Version = GetConfigValue("swagger.version")
	docs.SwaggerInfo.Host = GetConfigValue("swagger.host")
	docs.SwaggerInfo.BasePath = GetConfigValue("swagger.basePath")
	docs.SwaggerInfo.Schemes = []string{GetConfigValue("swagger.schema")}
}
