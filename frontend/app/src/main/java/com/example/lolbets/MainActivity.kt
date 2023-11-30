package com.example.lolbets

import android.R
import android.content.ContentValues.TAG
import android.content.IntentSender.SendIntentException
import android.os.Bundle
import android.util.Log
import androidx.activity.ComponentActivity
import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.compose.setContent
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Surface
import androidx.compose.runtime.getValue
import androidx.compose.ui.Modifier
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import androidx.lifecycle.lifecycleScope
import androidx.lifecycle.viewmodel.compose.viewModel
import com.example.lolbets.ui.sign_in.GoogleAuthUiClient
import com.example.lolbets.ui.sign_in.SignInViewModel
//import com.example.lolbets.network.KtorRealtimeMessagingClient
import com.example.lolbets.ui.theme.LolBetsTheme
import com.google.android.gms.auth.api.identity.GetSignInIntentRequest
import com.google.android.gms.auth.api.identity.Identity
import com.google.android.gms.auth.api.signin.GoogleSignIn
import com.google.android.gms.auth.api.signin.GoogleSignInClient
import com.google.android.gms.auth.api.signin.GoogleSignInOptions
import kotlinx.coroutines.launch


class MainActivity : ComponentActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            LolBetsTheme {
                // A surface container using the 'background' color from the theme
                Surface(
                    modifier = Modifier.fillMaxSize(),
                    color = MaterialTheme.colorScheme.background
                ) {
                    val googleAuthUiClient by lazy {
                        GoogleAuthUiClient(
                            context = applicationContext,
                            oneTapClient = Identity.getSignInClient(applicationContext)
                        )
                    }

                    LolBetsApp(googleAuthUiClient)
                }
            }
        }
    }
}