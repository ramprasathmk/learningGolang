# API Documentation

## Package: API

### Structures:
#### Employee
Represents an employee entity with the following fields:
- `Id` (int) - Employee ID
- `Name` (string) - Employee Name
- `Mobile` (string) - Employee Mobile Number
- `Email` (string) - Employee Email Address
- `Address` (string) - Employee Address
- `Dob` (string) - Employee Date of Birth

#### EmployeesList
Type alias for a slice of `Employee`.

#### Response
API response structure:
- `Status` (string) - API response status (`"S"` for success, `"E"` for error)
- `ErrMsg` (string) - Error message, if any
- `Employees` (`EmployeesList`) - List of employees

## API Endpoints

### `GET /get`
Retrieves all employee data.
- Response:
  - `Response` JSON containing status, employees list or error message.

### `POST /post`
Inserts a new employee.
- Request Body:
  - `Employee` JSON object
- Validation:
  - Employee ID must be unique.
- Response:
  - `Response` JSON indicating success or failure.

### `PUT /update`
Updates an existing employee's data.
- Request Body:
  - `Employee` JSON object (matching ID required)
- Validation:
  - Employee ID must exist.
- Response:
  - `Response` JSON indicating success or failure.

### `DELETE /delete`
Deletes an employee by ID.
- Request Body:
  - JSON object with `Id` field
- Validation:
  - Employee ID must exist.
- Response:
  - `Response` JSON indicating success or failure.

## Constants (common.go)
- `lStatusE` = "E" (Error status)
- `lStatusS` = "S" (Success status)
- `lErrMessage` = "something went wrong"
- `lEmpNotFoundErr` = "employee not found"
- `lEmpExistErr` = "employee already exist"

## Global Variables (common.go)
- `lEmpArr` - Predefined list of employees

## Server Configuration
- Listens on port `29100`
- Logs written to `./log/logfile<timestamp>.txt`
- Supports CORS with unrestricted access

