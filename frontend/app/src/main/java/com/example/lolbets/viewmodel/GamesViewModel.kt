package com.example.lolbets.viewmodel

import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.lolbets.model.Game
import com.example.lolbets.model.League
import com.example.lolbets.model.Team
import com.example.lolbets.network.LolApi
import kotlinx.coroutines.launch
import retrofit2.HttpException
import java.io.IOException
import java.text.SimpleDateFormat
import java.time.ZonedDateTime
import java.time.format.DateTimeFormatter
import java.util.Locale


sealed interface GameUiState {
    data class Success(val games: List<Game>) : GameUiState
    object Error : GameUiState
    object Loading : GameUiState
}

class GamesViewModel : ViewModel() {
    var gameUiState: GameUiState by mutableStateOf(GameUiState.Loading)
        private set


    init {
        getGames()
    }

    private fun formatTimestamp(timestamp: String): String {
        try {
            val inputFormat = SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss'Z'", Locale.getDefault())
            val outputFormat = SimpleDateFormat("dd/MM - HH:mm", Locale.getDefault())

            val date = inputFormat.parse(timestamp)
            return outputFormat.format(date!!)
        } catch (e: Exception) {
            // Handle parsing errors
            e.printStackTrace()
            return timestamp
        }
    }

    fun updateGame(game: Game) {
        // Loop games searching where game.Id is the same
        if (gameUiState is GameUiState.Success) {
            val currentState = gameUiState as GameUiState.Success
            val games = currentState.games.toMutableList()

            // Find the index of the game to update
            val gameIndex = games.indexOfFirst { it.id == game.id }

            // If the game is found, update it
            if (gameIndex != -1) {
                games[gameIndex] = game

                // Update the UI state with the new list of games
                gameUiState = GameUiState.Success(games)
            }
        }
    }

    fun getGames() {
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

                //Pillamos lista con los datos de los equipos que queremos
                val teamsList = LolApi.retrofitService.getTeams(unknownTeamsList.joinToString(","))

                //Pillamos lista con los datos de las ligas que queremos
                val leaguesList = LolApi.retrofitService.getLeagues(unknownLeaguesList.joinToString(","))

                //transformamos ambos en un map
                val teamsMap = teamsList.associateBy { it.name }
                val leaguesMap = leaguesList.associateBy { it.name }

                //Creamos la lista de games, es la que se mostrara por pantalla
                var finalGames = listOf<Game>()
                gamesList.forEach { game ->
                    val teamImage1 = teamsMap[game.team1]?.image
                    val teamImage2 = teamsMap[game.team2]?.image
                    val teamCode1 = teamsMap[game.team1]?.code
                    val teamCode2 = teamsMap[game.team2]?.code
                    val teamName1 = teamsMap[game.team1]?.name
                    val teamName2 = teamsMap[game.team2]?.name
                    val leagueImage = leaguesMap[game.league]?.image
                    val leagueRegion = leaguesMap[game.league]?.region

                    //Transform date to correct form
                    val gameDate = formatTimestamp(game.time)

                    finalGames = finalGames + Game(
                        game.id,
                        Team(teamName1!!, teamCode1!!, teamImage1!!),
                        Team(teamName2!!, teamCode2!!, teamImage2!!), League("0", game.league, leagueRegion!!, leagueImage!!),
                        gameDate, game.bets1,  game.bets2,
                        game.completed, game.blockName, game.strategy)

                }

                GameUiState.Success(
                    finalGames
                )
            } catch (e: IOException) {
                GameUiState.Error
            } catch (e: HttpException) {
                GameUiState.Error
            } catch (e: Exception) {
                GameUiState.Success(
                    listOf()
                )
            }
        }
    }
}
