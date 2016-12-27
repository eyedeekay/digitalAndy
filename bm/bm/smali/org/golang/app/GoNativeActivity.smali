.class public Lorg/golang/app/GoNativeActivity;
.super Landroid/app/NativeActivity;
.source "GoNativeActivity.java"


# static fields
.field private static goNativeActivity:Lorg/golang/app/GoNativeActivity;


# direct methods
.method public constructor <init>()V
    .locals 0

    .prologue
    .line 16
    invoke-direct {p0}, Landroid/app/NativeActivity;-><init>()V

    .line 17
    sput-object p0, Lorg/golang/app/GoNativeActivity;->goNativeActivity:Lorg/golang/app/GoNativeActivity;

    .line 18
    return-void
.end method

.method private load()V
    .locals 3

    .prologue
    .line 49
    :try_start_0
    invoke-virtual {p0}, Lorg/golang/app/GoNativeActivity;->getPackageManager()Landroid/content/pm/PackageManager;

    move-result-object v0

    invoke-virtual {p0}, Lorg/golang/app/GoNativeActivity;->getIntent()Landroid/content/Intent;

    move-result-object v1

    invoke-virtual {v1}, Landroid/content/Intent;->getComponent()Landroid/content/ComponentName;

    move-result-object v1

    const/16 v2, 0x80

    invoke-virtual {v0, v1, v2}, Landroid/content/pm/PackageManager;->getActivityInfo(Landroid/content/ComponentName;I)Landroid/content/pm/ActivityInfo;

    move-result-object v0

    .line 51
    iget-object v1, v0, Landroid/content/pm/ActivityInfo;->metaData:Landroid/os/Bundle;

    if-nez v1, :cond_0

    .line 52
    const-string v0, "Go"

    const-string v1, "loadLibrary: no manifest metadata found"

    invoke-static {v0, v1}, Landroid/util/Log;->e(Ljava/lang/String;Ljava/lang/String;)I

    .line 60
    :goto_0
    return-void

    .line 55
    :cond_0
    iget-object v0, v0, Landroid/content/pm/ActivityInfo;->metaData:Landroid/os/Bundle;

    const-string v1, "android.app.lib_name"

    invoke-virtual {v0, v1}, Landroid/os/Bundle;->getString(Ljava/lang/String;)Ljava/lang/String;

    move-result-object v0

    .line 56
    invoke-static {v0}, Ljava/lang/System;->loadLibrary(Ljava/lang/String;)V
    :try_end_0
    .catch Ljava/lang/Exception; {:try_start_0 .. :try_end_0} :catch_0

    goto :goto_0

    .line 57
    :catch_0
    move-exception v0

    .line 58
    const-string v1, "Go"

    const-string v2, "loadLibrary failed"

    invoke-static {v1, v2, v0}, Landroid/util/Log;->e(Ljava/lang/String;Ljava/lang/String;Ljava/lang/Throwable;)I

    goto :goto_0
.end method


# virtual methods
.method getRune(III)I
    .locals 4

    .prologue
    const/4 v0, -0x1

    .line 26
    :try_start_0
    invoke-static {p1}, Landroid/view/KeyCharacterMap;->load(I)Landroid/view/KeyCharacterMap;

    move-result-object v1

    invoke-virtual {v1, p2, p3}, Landroid/view/KeyCharacterMap;->get(II)I
    :try_end_0
    .catch Landroid/view/KeyCharacterMap$UnavailableException; {:try_start_0 .. :try_end_0} :catch_1
    .catch Ljava/lang/Exception; {:try_start_0 .. :try_end_0} :catch_0

    move-result v1

    .line 27
    if-nez v1, :cond_0

    .line 35
    :goto_0
    return v0

    :cond_0
    move v0, v1

    .line 30
    goto :goto_0

    .line 33
    :catch_0
    move-exception v1

    .line 34
    const-string v2, "Go"

    const-string v3, "exception reading KeyCharacterMap"

    invoke-static {v2, v3, v1}, Landroid/util/Log;->e(Ljava/lang/String;Ljava/lang/String;Ljava/lang/Throwable;)I

    goto :goto_0

    .line 31
    :catch_1
    move-exception v1

    goto :goto_0
.end method

.method getTmpdir()Ljava/lang/String;
    .locals 1

    .prologue
    .line 21
    invoke-virtual {p0}, Lorg/golang/app/GoNativeActivity;->getCacheDir()Ljava/io/File;

    move-result-object v0

    invoke-virtual {v0}, Ljava/io/File;->getAbsolutePath()Ljava/lang/String;

    move-result-object v0

    return-object v0
.end method

.method public onCreate(Landroid/os/Bundle;)V
    .locals 0

    .prologue
    .line 64
    invoke-direct {p0}, Lorg/golang/app/GoNativeActivity;->load()V

    .line 65
    invoke-super {p0, p1}, Landroid/app/NativeActivity;->onCreate(Landroid/os/Bundle;)V

    .line 66
    return-void
.end method
