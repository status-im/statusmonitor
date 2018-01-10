package main

import "testing"

func TestDumpSysOutput(t *testing.T) {
	dumpsys, err := NewDumpSysOutput(testDumpSysOutput)
	if err != nil {
		t.Fatalf("Expected err to be nil, but got %v", err)
	}
	expected := int64(10178)
	if dumpsys.UID != expected {
		t.Fatalf("Expected UID to be %v, but got %v", expected, dumpsys.UID)
	}
}

// sample real world output for testing
var testDumpSysOutput = `Activity Resolver Table:
  Non-Data Actions:
      android.intent.action.MAIN:
        f0569b im.status.ethereum/.MainActivity filter 5c6c43c
          Action: "android.intent.action.MAIN"
          Category: "android.intent.category.LAUNCHER"

Receiver Resolver Table:
  Non-Data Actions:
      android.net.conn.CONNECTIVITY_CHANGE:
        79a4a38 im.status.ethereum/com.instabug.library.network.InstabugNetworkReceiver filter 8715209
          Action: "android.net.conn.CONNECTIVITY_CHANGE"
      com.android.vending.INSTALL_REFERRER:
        e9eda11 im.status.ethereum/com.google.android.gms.measurement.AppMeasurementInstallReferrerReceiver filter b34dc0e
          Action: "com.android.vending.INSTALL_REFERRER"
      com.google.android.c2dm.intent.RECEIVE:
        a3a3776 im.status.ethereum/com.google.firebase.iid.FirebaseInstanceIdReceiver filter 162262f
          Action: "com.google.android.c2dm.intent.RECEIVE"
          Category: "im.status.ethereum"

Service Resolver Table:
  Non-Data Actions:
      com.google.firebase.INSTANCE_ID_EVENT:
        2f3c177 im.status.ethereum/com.evollu.react.fcm.InstanceIdService filter 72625c2
          Action: "com.google.firebase.INSTANCE_ID_EVENT"
        fc5d1e4 im.status.ethereum/com.google.firebase.iid.FirebaseInstanceIdService filter cb5fd3
          Action: "com.google.firebase.INSTANCE_ID_EVENT"
          mPriority=-500, mHasPartialTypes=false
      com.google.firebase.MESSAGING_EVENT:
        786064d im.status.ethereum/com.evollu.react.fcm.MessagingService filter 85fa20d
          Action: "com.google.firebase.MESSAGING_EVENT"
        68a5502 im.status.ethereum/com.google.firebase.messaging.FirebaseMessagingService filter dfc4710
          Action: "com.google.firebase.MESSAGING_EVENT"
          mPriority=-500, mHasPartialTypes=false

Permissions:
  Permission [im.status.ethereum.permission.C2D_MESSAGE] (58fe13):
    sourcePackage=im.status.ethereum
    uid=10178 gids=null type=0 prot=signature
    perm=Permission{9002850 im.status.ethereum.permission.C2D_MESSAGE}
    packageSetting=PackageSetting{f39ba49 im.status.ethereum/10178}

Registered ContentProviders:
  im.status.ethereum/com.google.firebase.provider.FirebaseInitProvider:
    Provider{fccbf4e im.status.ethereum/com.google.firebase.provider.FirebaseInitProvider}
  im.status.ethereum/android.support.v4.content.FileProvider:
    Provider{cb686f im.status.ethereum/android.support.v4.content.FileProvider}

ContentProvider Authorities:
  [im.status.ethereum.firebaseinitprovider]:
    Provider{fccbf4e im.status.ethereum/com.google.firebase.provider.FirebaseInitProvider}
      applicationInfo=ApplicationInfo{3b84e77 im.status.ethereum}
  [im.status.ethereum.provider]:
    Provider{cb686f im.status.ethereum/android.support.v4.content.FileProvider}
      applicationInfo=ApplicationInfo{3b84e77 im.status.ethereum}

Key Set Manager:
  [im.status.ethereum]
      Signing KeySets: 90

Packages:
  Package [im.status.ethereum] (f39ba49):
    userId=10178
    pkg=Package{547797c im.status.ethereum}
    codePath=/data/app/im.status.ethereum-qu6Tk7NWOu21X1TeL7giAQ==
    resourcePath=/data/app/im.status.ethereum-qu6Tk7NWOu21X1TeL7giAQ==
    legacyNativeLibraryDir=/data/app/im.status.ethereum-qu6Tk7NWOu21X1TeL7giAQ==/lib
    primaryCpuAbi=armeabi-v7a
    secondaryCpuAbi=null
    versionCode=2054 minSdk=18 targetSdk=23
    versionName=0.9.10-497-gc531ece0+
    splits=[base]
    apkSigningVersion=2
    applicationInfo=ApplicationInfo{3b84e77 im.status.ethereum}
    flags=[ HAS_CODE ALLOW_CLEAR_USER_DATA ALLOW_BACKUP LARGE_HEAP ]
    dataDir=/data/user/0/im.status.ethereum
    supportsScreens=[small, medium, large, xlarge, resizeable, anyDensity]
    timeStamp=2018-01-03 15:00:31
    firstInstallTime=2018-01-03 15:00:36
    lastUpdateTime=2018-01-03 15:00:36
    signatures=PackageSignatures{85ab205 [f4d5fb87]}
    installPermissionsFixed=true installStatus=1
    pkgFlags=[ HAS_CODE ALLOW_CLEAR_USER_DATA ALLOW_BACKUP LARGE_HEAP ]
    declared permissions:
      im.status.ethereum.permission.C2D_MESSAGE: prot=signature, INSTALLED
    requested permissions:
      android.permission.INTERNET
      android.permission.NFC
      android.permission.ACCESS_NETWORK_STATE
      android.permission.ACCESS_WIFI_STATE
      android.permission.READ_PROFILE
      android.permission.CAMERA
      android.permission.READ_EXTERNAL_STORAGE
      android.permission.WRITE_EXTERNAL_STORAGE
      android.permission.READ_CONTACTS
      android.permission.RECEIVE_SMS
      android.permission.ACCESS_FINE_LOCATION
      android.permission.ACCESS_COARSE_LOCATION
      com.google.android.c2dm.permission.RECEIVE
      android.permission.WAKE_LOCK
      im.status.ethereum.permission.C2D_MESSAGE
      com.sec.android.provider.badge.permission.READ
      com.sec.android.provider.badge.permission.WRITE
      com.htc.launcher.permission.READ_SETTINGS
      com.htc.launcher.permission.UPDATE_SHORTCUT
      com.sonyericsson.home.permission.BROADCAST_BADGE
      com.sonymobile.home.permission.PROVIDER_INSERT_BADGE
      com.anddoes.launcher.permission.UPDATE_COUNT
      com.majeur.launcher.permission.UPDATE_BADGE
      com.huawei.android.launcher.permission.CHANGE_BADGE
      com.huawei.android.launcher.permission.READ_SETTINGS
      com.huawei.android.launcher.permission.WRITE_SETTINGS
      android.permission.READ_APP_BADGE
      com.oppo.launcher.permission.READ_SETTINGS
      com.oppo.launcher.permission.WRITE_SETTINGS
      me.everything.badger.permission.BADGE_COUNT_READ
      me.everything.badger.permission.BADGE_COUNT_WRITE
    install permissions:
      com.google.android.c2dm.permission.RECEIVE: granted=true
      android.permission.NFC: granted=true
      android.permission.READ_PROFILE: granted=true
      android.permission.INTERNET: granted=true
      android.permission.ACCESS_NETWORK_STATE: granted=true
      im.status.ethereum.permission.C2D_MESSAGE: granted=true
      android.permission.ACCESS_WIFI_STATE: granted=true
      android.permission.WAKE_LOCK: granted=true
    User 0: ceDataInode=395674 installed=true hidden=false suspended=false stopped=false notLaunched=false enabled=0 instant=false virtual=false
    overlay paths:
      /vendor/overlay/framework-res__auto_generated_rro.apk
      /vendor/overlay/Pixel/PixelThemeOverlay.apk
      gids=[3003]
      runtime permissions:

Package Changes:
  Sequence number=74
  User 0:
    seq=17, package=im.status.ethereum
    seq=26, package=com.twitter.android
    seq=27, package=com.google.android.play.games
    seq=31, package=com.android.stk
    seq=32, package=com.google.android.ims
    seq=37, package=com.google.android.gms
    seq=39, package=com.google.android.apps.maps
    seq=40, package=com.vrem.wifianalyzer
    seq=43, package=com.asus.filemanager
    seq=44, package=com.VaRs.VRPlayerPRO
    seq=50, package=com.google.android.apps.docs.editors.docs
    seq=51, package=com.google.android.tts
    seq=52, package=com.google.android.apps.tachyon
    seq=53, package=com.spotify.music
    seq=66, package=com.instagram.android
    seq=67, package=com.google.android.apps.turbo
    seq=68, package=com.google.android.keep
    seq=69, package=com.facebook.lite
    seq=71, package=com.google.android.apps.messaging
    seq=72, package=com.google.android.youtube
    seq=73, package=com.google.android.calendar


Dexopt state:
  [im.status.ethereum]
    path: /data/app/im.status.ethereum-qu6Tk7NWOu21X1TeL7giAQ==/base.apk
      arm: /data/app/im.status.ethereum-qu6Tk7NWOu21X1TeL7giAQ==/oat/arm/base.odex[status=kOatUpToDate, compilation_filt
      er=speed-profile]


Compiler stats:
  [im.status.ethereum]
     base.apk - 4586

 `
