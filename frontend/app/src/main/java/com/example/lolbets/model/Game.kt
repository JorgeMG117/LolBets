package com.example.lolbets.model

import androidx.annotation.DrawableRes
import androidx.annotation.StringRes
import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class GameApi(
    val id: Int,
    val team1: String,
    val team2: String,
    val league: String,
    val time: String,
    val bets1: Int,
    val bets2: Int,
    val completed: Int,
    val blockName: String,
    val strategy: String,
)

data class Game(
    val team1: Team,
    val team2: Team,
    val league: League,
    val date: String,
    val betsTeam1: Int,
    val betsTeam2: Int,
)

data class Todo(
    var userId: Int,
    var id: Int,
    var title: String,
    var completed: Boolean
)

const val BASE_URL = "https://jsonplaceholder.typicode.com/"
/*
interface APIService {
    @GET("todos")
    suspend fun getTodos(): List<Todo>

    companion object {
        var apiService: APIService? = null
        fun getInstance(): APIService {
            if (apiService == null) {
                apiService = Retrofit.Builder()
                    .baseUrl(BASE_URL)
                    .addConverterFactory(GsonConverterFactory.create())
                    .build().create(APIService::class.java)
            }
            return apiService!!
        }
    }
}*/