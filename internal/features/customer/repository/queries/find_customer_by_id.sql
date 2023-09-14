
SELECT
    c.id, c.first_name, c.last_name, c.email, c.password,
    a.id "address.id", a.street_name "address.street_name", a.house_number "address.house_number",
    a.city "address.city", a.zip_code "address.zip_code", a.country "address.country",
    a.state "address.state", a.region "address.region", a.extra_info "address.extra_info",
    a.created_at "address.created_at", a.updated_at "address.updated_at",
    c.created_at, c.updated_at
FROM customers c
    INNER JOIN addresses a ON c.address_id = a.id
WHERE c.id = $1