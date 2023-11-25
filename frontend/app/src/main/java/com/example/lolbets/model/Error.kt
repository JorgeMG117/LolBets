package com.example.lolbets.model

import kotlinx.serialization.Serializable

@Serializable
data class ErrorResponse(
    val error: String
)