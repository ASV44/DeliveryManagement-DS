package proxy

const (
	GetRequestToRoot      = "GET request on '/' route"
	WelcomeMessage        = "Delivery Management Distributed System ~DS Proxy is on the fly\n"
	TimeMessage           = "Proxy Time: "
	ForwardRequest        = "Forward request with length : "
	ForwardResponse       = "Forward response with length : "
	ForwardOrderPost      = "Forwarding POST of order : "
	ForwardOrdersRegister = "Forwarding POST of multiple orders : "
	ForwardOrderUpdate    = "Forwarding PUT request for update order with ID %s : "
	ForwardOrderDelete    = "Forwarding DELETE of order with ID %s : "
	ForwardOrdersDelete   = "Forwarding DELETE of multiple orders : "
	ForwardOrderGetById   = "Forwarding GET of order with ID %s : "
	ForwardOrdersGetByAwb = "Forwarding GET of orders with AWB number %s : "
	ForwardOrdersGet      = "Forwarding GET of all orders : "
	DataNotFoundInCache   = "Not Found data in cache : "
	AddDataToCache        = "Data Added to cache :"
)
