package offer_routes

/*
GET /feed - List all services with basic filters (like type or location - default if not provided).
POST /feed/search - Perform a full-text search on services. The POST method
*/

/*
PUT /discovery/index - (Re-)index services for search. This might be an admin or system-only endpoint.
DELETE /discovery/index - Delete a specific index. Again, this could be restricted to admin or system use.
POST /discovery/suggest - Get autocomplete suggestions for a search term, leveraging Elasticsearch's capabilities.
*/
