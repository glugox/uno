package server

func ConfigureControllers(srv *Server) {

	/*for _, r := range srv.Routes {
		srv.Mux.Handle(r.Regex, r.Handler)
	}*/

	// // Menu Controller
	// menuMux := http.NewServeMux()
	// menuMux.HandleFunc("/", MenuController(srv.App))
	// srv.Mux.Handle("/menu", menuMux)

	// // // Menu Items Controller
	// menuItemsMux := http.NewServeMux()
	// menuItemsMux.HandleFunc("/", MenuItemsController(srv.App))
	// srv.Mux.Handle("/menu-items", menuItemsMux)

}
