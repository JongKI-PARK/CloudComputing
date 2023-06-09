package com.example.demo.controller;

import com.example.demo.controller.dto.ResponseWrapper;
import com.example.demo.service.EnrollmentPlanService;
import com.example.demo.service.dto.EnrollmentCreateRequestDto;
import lombok.AllArgsConstructor;
import org.springframework.web.bind.annotation.*;


@RestController
@AllArgsConstructor
public class EnrollmentPlanController {
    private final EnrollmentPlanService enrollmentPlanService;

    @GetMapping("/planner")
    public ResponseWrapper getStudentEnrollmentPlan(@RequestParam Long studentId) {
        return new ResponseWrapper("success", enrollmentPlanService.getStudentEnrollmentPlan(studentId));
    }

    @PostMapping("/planner")
    public ResponseWrapper saveEnrollmentPlan(@RequestBody EnrollmentCreateRequestDto requestDto) {
        return new ResponseWrapper("success", enrollmentPlanService.saveEnrollmentPlan(requestDto));
    }
}
