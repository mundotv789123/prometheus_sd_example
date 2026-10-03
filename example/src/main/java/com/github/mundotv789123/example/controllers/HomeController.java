package com.github.mundotv789123.example.controllers;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/")
public class HomeController {

    @GetMapping
    public ResponseEntity<String> index() {
        return ResponseEntity.ok("Olá mundo");
    }
    
    @GetMapping("/erro")
    public ResponseEntity<String> erro() {
        String texto = null;
        texto.equals("erro 500 proposital");
        return ResponseEntity.ok("Olá mundo");
    }
}
