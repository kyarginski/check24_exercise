# API documentation

## Requirements

- Creating or inserting blog entries
- Creating or inserting comments on blog entries
- Read previously saved data records (blog entries)
  - The following filters are possible 
    - Sorting
    - Number
    - Text length
  - The following information is displayed in the entry overview
    - Date,
    - author,
    - URL to the image,
    - text (see filter above) 
    - and number of comments
    
- Output of a complete entry with
  - Date,
  - text,
  - author,
  - comment(s),
  - URL to the image
  
- Updating existing data records

- Individual deletion of blog entries

## Description

### Creating Blog Entries
    Method: POST
    Path: /api/blog/entries
    Description: Create a new blog entry. The request body should include data about the entry, such as title, text, author, and an optional image link.
    Response: If the creation is successful, return a 201 (Created) status code and information about the created entry.

### Creating Comments on Blog Entries
    Method: POST
    Path: /api/blog/entries/{entry_id}/comments
    Description: Create a new comment for a specific blog entry (where {entry_id} is the identifier of the entry). The request body should include data about the comment, such as name, text, and optionally an email and URL.
    Response: If the creation is successful, return a 201 (Created) status code and information about the created comment.

### Reading Previously Saved Data Records (Blog Entries)
    Method: GET
    Path: /api/blog/entries
    Description: Retrieve a list of blog entries with options for filtering, sorting, and limiting the number of entries. Use query parameters for filtering (e.g., by text), sorting, and limiting the number of entries.
    Response: Return a list of blog entries based on the specified filters, sorting, and limiting, along with information about each entry.

### Output of a Complete Entry
    Method: GET
    Path: /api/blog/entries/{entry_id}
    Description: Retrieve complete information about a specific blog entry (where {entry_id} is the entry's identifier).
    Response: Return detailed information about the entry, including the date, text, author, comments, image link, and more.

### Updating Existing Data Records
    Method: PUT
    Path: /api/blog/entries/{entry_id}
    Description: Update an existing blog entry (where {entry_id} is the entry's identifier). The request body should include new data for the update.
    Response: If the update is successful, return a 200 (OK) status code and information about the updated entry.

### Individual Deletion of Blog Entries
    Method: DELETE
    Path: /api/blog/entries/{entry_id}
    Description: Delete a specific blog entry (where {entry_id} is the entry's identifier).
    Response: If the deletion is successful, return a 204 (No Content) status code without a response body.
