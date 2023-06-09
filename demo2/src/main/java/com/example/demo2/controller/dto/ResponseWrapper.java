package com.example.demo2.controller.dto;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class ResponseWrapper {
    private String status;
    private Object data;
}
