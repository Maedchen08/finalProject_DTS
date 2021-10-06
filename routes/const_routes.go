package routes

//role
const POST_ROLE = "/role"
const GETALL_ROLE = "/role"
const GETBYID_ROLE = "/role/:id"

//agents
const POST_AGENTS = "/agents"
const GETALL_AGENTS = "/agents"
const GETBYID_AGENTS = "/agents/:id"

//customer
const POST_CUSTOMER = "/customer"
const GETALL_CUSTOMER = "/customer"
const GETBYID_CUSTOMER = "/customer/:id"

//serviceTransaction
const POST_SERVICE_TRANSACTION = "/st"
const GETALL_SERVICE_TRANSACTION = "/st"
const GETBYID_SERVICE_TRANSACTION = "/st/:id"

//typeServiceTransaction
const POST_TYPE_SERVICE_TRANSACTION = "/ts"
const GETALL_TYPE_SERVICE_TRANSACTION = "/ts"
const GETBYID_TYPE_SERVICE_TRANSACTION = "/ts/:id"

//transaction
const POST_TRANSACTION = "/transaksi/create"
const POST_AGENT_TRANSACTION = "/transaksi/cariagent"
const POST_CONFIRM_TRANSACTION = "/transaksi/dikonfirmasi"
const POST_REJECT_TRANSACTION = "/transaksi/dibatalkan"
const POST_DONE_TRANSACTION = "/transaksi/selesai"
const GETBYID_TRANSACTION = "/transaksi/:id"
const GETBYID_TRANSACTION_CUSTOMER = "/transaksi/customer/:id"
const GETBYID_TRANSACTION_AGENT = "/transaksi/agent/:id"
const GETALL_TRANSACTION = "/transaksi"
const DELETE_TRANSACTION = "/transaksi/:id"


//auth
const POST_REGISTER = "/register"
const GETALL_USERS = "/users"
const LOGIN = "/login"
const GETBYID_USERS = "/users/:id"
