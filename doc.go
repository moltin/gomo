/*
Package gomo provides an API client for Moltin's headless ecommerce API.

Sign up for an account at https://moltin.com.

By Andrew Waters and the team at Moltin. Contributions are most welcome.

Getting Started

Create a client object, and authenticate:

	client := gomo.NewClient()
	err := client.Authenticate()

By default the environment variables MOLTIN_CLIENT_ID and MOLTIN_CLIENT_SECRET
will be used for authentication.

Requests are made to the API via the Get(), Post(), Put() and Delete()
functions, each of which accepts a path and a number of resources. These
resources can set the target for the returned data or the body of the request,
as will as configuring various other request options.

Requests and their responses are usually marshalled into core types
representing the various objects in the Moltin API.

So to read category ID c82d2f00-bc66-4c7d-984a-8765222abb98, along with its
associated products one might use:

	id := "c82d2f00-bc66-4c7d-984a-8765222abb98"
	var category core.Category
	var included struct {
		Products []core.Product `json:"products"`
	}
	err := client.Get(
		"categories/"+id,
		gomo.Include("products"),
		gomo.Included(&included),
		gomo.Data(&category)
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Category %s has products:\n", category.Name)
	for _, product := range included.Products {
		fmt.Println(product.ID)
	}

Flows

Flows allow arbitrary data to be associated with objects. To cope with them
in Go, create a struct for the flow fields and embed the core object to access
the core fields, eg:

	type MyProduct struct {
		core.Product
		Stars int `json:"stars"`
	}
	var product MyProduct
	err = client.Get(
		"/products/8610c22b-f3a5-48a6-a680-8e8902a74aac",
		gomo.Data(&product),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(
		"Product %s has %d stars\n",
		product.Product.Name,
		product.Stars,
	)


*/
package gomo
