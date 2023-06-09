package com.example.demo2.controller.dto;

import com.example.demo2.service.EnrollmentProcess;
import com.example.demo2.service.dto.EnrollmentCompleteDto;
import com.example.demo2.service.dto.EnrollmentOrderRequest;
import lombok.AllArgsConstructor;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
@AllArgsConstructor
public class EnrollmentOrderController {
    private final EnrollmentProcess enrollmentProcess;

    @PostMapping("/order")
    public ResponseWrapper saveEnrollmentComplete(@RequestBody EnrollmentOrderRequest orderRequest) {
        EnrollmentCompleteDto completeDto = enrollmentProcess.saveEnrollmentComplete(orderRequest);
        return new ResponseWrapper(completeDto == null ? "fail" : "success", completeDto);
    }

}
