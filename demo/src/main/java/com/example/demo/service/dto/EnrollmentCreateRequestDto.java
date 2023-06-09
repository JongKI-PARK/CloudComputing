package com.example.demo.service.dto;

import com.example.demo.entity.EnrollmentPlan;
import lombok.Getter;

@Getter
public class EnrollmentCreateRequestDto {
    private Long studentId;
    private Long subjectId;

    public EnrollmentPlan convertToEntity() {
        return new EnrollmentPlan(studentId, subjectId);
    }
}
