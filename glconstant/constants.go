package glconstant

const EMPTY_VALUE string = ""
const NULL_REF_VALUE_FOR_LONG int64 = -99
const DATETIME_FORMAT string = "20060102150405"
const DATE_FORMAT string = "20060102"
const YES string = "Y"
const NO string = "N"
const USER_SUPERADMIN int64 = -1
const ROLE_SUPERADMIN int64 = -1
const TENANT_SUPERADMIN int64 = -1

const ENV_LOG_FILE string = "LOG_FILE"
const ENV_LOG_LEVEL string = "LOG_LEVEL"
const ENV_DB_NAME string = "DB_NAME"
const ENV_DB_HOST string = "DB_HOST"
const ENV_DB_PORT string = "DB_PORT"
const ENV_DB_PASSWORD string = "DB_PASSWORD"
const ENV_DB_USER string = "DB_USER"

const LOG_FILE_DEFAULT string = "log/goleaf.log"
const LOG_LEVEL_DEFAULT string = "debug"

const HEADER_ROLE_ID string = "X-Role-Id"
const HEADER_TENANT_ID string = "X-Tenant-Id"
const HEADER_TENANT_SCHEMA string = "X-Tenant"
const HEADER_SESSION_ID string = "X-Session-Id"
const HEADER_JWT_TASK string = "X-Jwt-Task"
const HEADER_DATETIME string = "X-Datetime"

const DEFAULT_TIMEZONE string = "+0700"

const STATUS_OK string = "OK"
const STATUS_FAIL string = "FAIL"
