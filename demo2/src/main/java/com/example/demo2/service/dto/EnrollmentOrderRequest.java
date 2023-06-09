package com.example.demo2.service.dto;

import com.example.demo2.entity.EnrollmentComplete;
import lombok.Getter;

@Getter
public class EnrollmentOrderRequest {
    private Long studentId;
    private Long subjectId;

    public EnrollmentComplete convertToEnrollmentComplete() {
        return new EnrollmentComplete(studentId, subjectId);
    }
}
