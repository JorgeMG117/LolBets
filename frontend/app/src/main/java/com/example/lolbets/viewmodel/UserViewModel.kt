package com.example.lolbets.viewmodel

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.lolbets.R
import com.example.lolbets.data.BetUiState
import com.example.lolbets.model.Game
import com.example.lolbets.model.User
import com.example.lolbets.model.UserRequestBody
import com.example.lolbets.network.LolApi
import com.example.lolbets.ui.sign_in.UserData
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.launch

class UserViewModel : ViewModel() {

    private val _userState = MutableStateFlow(User(0, "R.drawable.koi", 1000, "Jiasf", "https://yt3.ggpht.com/ytc/APkrFKZKkpVjHnMHEb-DJNkgiz8lbI4Xmhq9-xnqDAXK_bg=s600-c-k-c0x00ffffff-no-rj-rp-mo"))
    val userState: StateFlow<User> = _userState.asStateFlow()

    fun placeBet(value: Int) {
        // TODO Check if value is negative

        _userState.update { currentState ->
            currentState.copy(
                coins = currentState.coins - value
            )
        }
    }

    fun getUserInfo(userData: UserData?, onUserUpdated: (Int) -> Unit) {
        println(userData)
        val requestBody = UserRequestBody(userData!!.userId)
        viewModelScope.launch {
            val userInfo = LolApi.retrofitService.postUser(requestBody)
            println(userInfo)
            //Update user view model with info
            //userInfo, userData
            _userState.update { currentState ->
                currentState.copy(
                    id = userInfo.id,
                    googleId = userInfo.name,
                    coins = userInfo.coins,
                    username = userData.username,
                    profilePictureUrl = userData.profilePictureUrl,
                )
            }
            onUserUpdated(userInfo.id)
        }
    }
}