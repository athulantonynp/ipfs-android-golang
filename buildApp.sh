cd go/ipfs
gomobile bind -v  -target=android -androidapi 24
cd -
rm -rf app/libs/ipfs.aar
if [ -f go/ipfs/ipfs.aar ]; then
    cp go/ipfs/ipfs.aar app/libs/
else
    echo "go/ipfs/ipfs.aar does not exist."
fi
./gradlew clean
./gradlew installDebug
adb shell am start -n "com.dreamcatcher.android/com.dreamcatcher.android.MainActivity"