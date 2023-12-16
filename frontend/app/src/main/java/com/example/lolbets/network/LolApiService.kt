package com.example.lolbets.network

import com.example.lolbets.model.ActiveBets
import com.example.lolbets.model.ErrorResponse
import com.example.lolbets.model.GameApi
import com.example.lolbets.model.League
import com.example.lolbets.model.Team
import com.example.lolbets.model.User
import com.example.lolbets.model.UserApi
import com.example.lolbets.model.UserRequestBody
import retrofit2.Retrofit
import com.jakewharton.retrofit2.converter.kotlinx.serialization.asConverterFactory
import kotlinx.serialization.json.Json
import okhttp3.MediaType.Companion.toMediaType
import retrofit2.Response
import retrofit2.http.Body
import retrofit2.http.GET
import retrofit2.http.POST
import retrofit2.http.Query

private const val BASE_URL = "https://lolbets-image-y5oknlrakq-no.a.run.app:443"
//https://lolbets-image-y5oknlrakq-no.a.run.app
//http://10.0.2.2:8080
//https://lolesports.com/schedule?leagues=lec,emea_masters,superliga,cblol-brazil,lck,lcl,lco,lcs,ljl-japan,lla,lpl,pcs,turkiye-sampiyonluk-ligi,vcs,wqs,cblol_academy,esports_balkan_league,elite_series,greek_legends,hitpoint_masters,lck_challengers_league,north_american_challenger_league,lcs_challengers_qualifiers,lfl,liga_portuguesa,nlc,pg_nationals,primeleague,ultraliga,arabian_league,tft_esports
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

    @GET("leagues")
    suspend fun getLeagues(@Query("leagues") leaguesList: String): List<League>

    @GET("teams")
    suspend fun getTeams(@Query("teams") teamsList: String): List<Team>

    //TODO Ojo si lo devuelve vacio
    @GET("activeBets")
    suspend fun getActiveBets(@Query("userId") userId: Int): List<ActiveBets>

    @POST("users")
    suspend fun postUser(@Body requestBody: UserRequestBody): UserApi
}

/**
 * A public Api object that exposes the lazy-initialized Retrofit service
 */
object LolApi {
    val retrofitService: LolApiService by lazy {
        retrofit.create(LolApiService::class.java)
    }
}