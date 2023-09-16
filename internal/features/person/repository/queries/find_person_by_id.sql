
SELECT
    p.id, p.first_name, p.last_name, p.email, p.password,
    a.id "address.id", a.street_name "address.street_name", a.house_number "address.house_number",
    a.city "address.city", a.zip_code "address.zip_code", a.country "address.country",
    a.state "address.state", a.region "address.region", a.extra_info "address.extra_info",
    a.created_at "address.created_at", a.updated_at "address.updated_at",
    p.created_at, p.updated_at
FROM persons p
    INNER JOIN addresses a ON p.address_id = a.id
WHERE p.id = $1