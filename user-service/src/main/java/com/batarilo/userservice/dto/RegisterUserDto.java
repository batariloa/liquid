package com.batarilo.userservice.dto;


import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;
import lombok.Setter;

@Builder
@AllArgsConstructor
@Getter
@Setter
public class RegisterUserDto {
    private String email;
    private String password;
    private String fullName;
}
