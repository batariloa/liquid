package com.batarilo.userservice.dto;


import jakarta.validation.constraints.Min;
import jakarta.validation.constraints.NotEmpty;
import jakarta.validation.constraints.NotNull;
import jakarta.validation.constraints.Size;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;
import lombok.Setter;



@Builder
@AllArgsConstructor
@Getter
@Setter
public class RegisterUserDto {
    @NotNull
    @NotEmpty
    private String email;
    @NotNull
    @NotEmpty
    @Size(min = 6, message = "Password must be at least 6 characters long.")
    private String password;
    @NotNull
    @NotEmpty
    private String fullName;
}
