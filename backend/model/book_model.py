from pydantic import BaseModel, Field, field_validator


class Book(BaseModel):
    id: int 
    title: str = Field(..., min_length=1)
    author: str
    year: int
    isbn: str

    @field_validator("year")
    def year_validator(cls,value):
        if value<1000 or value>2026:
            raise ValueError("Year must be between 1000 and 2026")
        return value
    @field_validator("isbn")
    def isbn_validator(cls,value):
        if len(value)!=13 and len(value)!=10:
            raise ValueError("ISBN must be 13 or 10 characters long")
        return value