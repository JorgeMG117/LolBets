package com.example.lolbets.network

import com.example.lolbets.model.Game
import com.example.lolbets.model.GameApi
import retrofit2.Retrofit
import com.jakewharton.retrofit2.converter.kotlinx.serialization.asConverterFactory
import kotlinx.serialization.json.Json
import okhttp3.MediaType.Companion.toMediaType
import retrofit2.http.GET

private const val BASE_URL = "http://10.0.2.2:8080"
    //"https://android-kotlin-fun-mars-server.appspot.com"

/**
 * Use the Retrofit builder to build a retrofit object using a kotlinx.serialization converter
 */
private val retrofit = Retrofit.Builder()
    .addConverterFactory(Json.asConverterFactory("application/json".toMediaType()))
    .baseUrl(BASE_URL)
    .build()

/**
 * Retrofit service object for creating api calls
 */
interface LolApiService {
    @GET("games")
    suspend fun getGames(): List<GameApi>
}

/**
 * A public Api object that exposes the lazy-initialized Retrofit service
 */
object LolApi {
    val retrofitService: LolApiService by lazy {
        retrofit.create(LolApiService::class.java)
    }
}