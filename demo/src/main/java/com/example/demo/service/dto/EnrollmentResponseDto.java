package com.example.demo.service.dto;

import com.example.demo.entity.EnrollmentPlan;
import lombok.AllArgsConstructor;
import lombok.Getter;

@Getter
@AllArgsConstructor
public class EnrollmentResponseDto {
    private Long id;
    private Long studentId;
    private Long subjectId;

    public EnrollmentResponseDto(EnrollmentPlan enrollmentPlan) {
        this.id = enrollmentPlan.getId();
        this.studentId = enrollmentPlan.getStudentId();
        this.subjectId = enrollmentPlan.getSubjectId();
    }
}
