package com.example.lolbets.ui

import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.lolbets.R
import com.example.lolbets.model.Game
import com.example.lolbets.model.League
import com.example.lolbets.model.Team
import com.example.lolbets.network.LolApi
import kotlinx.coroutines.launch
import retrofit2.HttpException
import java.io.IOException


sealed interface GameUiState {
    data class Success(val games: List<Game>) : GameUiState
    object Error : GameUiState
    object Loading : GameUiState
}

class GamesViewModel : ViewModel() {
    var gameUiState: GameUiState by mutableStateOf(GameUiState.Loading)
        private set


    init {
        getMarsPhotos()
    }

    fun getMarsPhotos() {
        viewModelScope.launch {
            gameUiState = GameUiState.Loading
            gameUiState = try {
                val gamesList = LolApi.retrofitService.getGames()

                //Para cada partido
                //Si team1.imagen no esta en cache o no se ha añadido ya
                    //Añadir a lista de peticion de equipo para imagen
                //Lo mismo con team2 y liga

                //Hacer peticion de teams
                //Hacer peticion de ligas
                var unknownTeamsList = listOf<String>()
                var unknownLeaguesList = listOf<String>()
                gamesList.forEach { game ->
                    if (!unknownTeamsList.contains(game.team1)) {
                        unknownTeamsList = unknownTeamsList + game.team1
                    }
                    if (!unknownTeamsList.contains(game.team2)) {
                        unknownTeamsList = unknownTeamsList + game.team2
                    }

                    if (!unknownLeaguesList.contains(game.league)) {
                        unknownLeaguesList = unknownLeaguesList + game.league
                    }

                }

                println(unknownTeamsList)

                //Pillamos lista con los datos de los equipos que queremos
                val teamsList = LolApi.retrofitService.getTeams(unknownTeamsList.joinToString(","))
                println(teamsList)

                //Pillamos lista con los datos de las ligas que queremos
                //val leaguesList = LolApi.retrofitService.getLeagues(unknownLeaguesList)

                //transformamos ambos en un map
                val teamsMap = teamsList.associateBy { it.name }
                //val leaguesMap = leaguesList.associateBy { it.name }

                //Creamos la lista de games, es la que se mostrara por pantalla
                var finalGames = listOf<Game>()
                gamesList.forEach { game ->
                    val teamImage1 = teamsMap[game.team1]?.image
                    val teamImage2 = teamsMap[game.team2]?.image
                    val teamCode1 = teamsMap[game.team1]?.code
                    val teamCode2 = teamsMap[game.team1]?.code
                    finalGames = finalGames + Game(
                        Team(teamCode1!!, teamImage1!!), Team(
                            teamCode2!!, teamImage2!!), League(R.string.league_name_lec, R.drawable.lec), game.time, game.bets1,  game.bets2)

                }

                GameUiState.Success(
                    finalGames
                )
            } catch (e: IOException) {
                GameUiState.Error
            } catch (e: HttpException) {
                GameUiState.Error
            }
        }
    }
}
