package com.batarilo.userservice.controller;

import com.batarilo.userservice.dto.LoginResponseDto;
import com.batarilo.userservice.dto.LoginUserDto;
import com.batarilo.userservice.dto.RegisterUserDto;
import com.batarilo.userservice.model.User;
import com.batarilo.userservice.service.AuthenticationService;
import com.batarilo.userservice.service.JwtService;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.apache.coyote.BadRequestException;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/v1/auth")
@RequiredArgsConstructor
@Validated
public class AuthenticationController
{
    private final AuthenticationService authenticationService;
    private final JwtService jwtService;

    @PostMapping("/register")
    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<User> register(@RequestBody
                                             @Valid RegisterUserDto registerUserDto) throws BadRequestException
    {
        if(authenticationService.isEmailInUse(registerUserDto.getEmail())) {
            throw new BadRequestException("Email already in use.");
        }

        User registeredUser = authenticationService.signup(registerUserDto);

        return ResponseEntity.ok(registeredUser);
    }

    @PostMapping("/login")
    public ResponseEntity<LoginResponseDto> login(@RequestBody LoginUserDto loginUserDto) {

        User authenticatedUser = authenticationService.authenticate(loginUserDto);

        String jwtToken = jwtService.generateToken(authenticatedUser);

        LoginResponseDto loginResponse =
            LoginResponseDto.builder()
                .token(jwtToken)
                .expiresIn(jwtService.getExpirationTime())
                .build();

        return ResponseEntity.ok(loginResponse);
    }
}
