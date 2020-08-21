
# Oculus Rift CV1 传感器追踪状态检查工具  
# (Oculus Rift CV1 Sensor Tracking Status Checking Tool)  
这是一个可以显示单独的Oculus传感器追踪状态的工具，使用web显示，可以轻松地导入OBS  
This is a tool that can show you individual oculus sensor tracking status, output as a web, easy cooperate with OBS

### 初始化/initialization
> [Releases](https://github.com/VAIO-Dong/Oculus-Rift-CV1-Sensor-Tracking-Status-Checking-Tool/releases)  
你可以从上面的Releases中下载到我们打包好的整个程序  
找个你喜欢的地方建个文件夹把他放进去 双击main.exe启动  
输入传感器1的序列号然后回车（序列号可以在Oculus的应用程序里看）  
输入传感器2的序列号然后回车  
输入EXIT回车（3/4传感器的web端还没做XD）  
然后找到在他的旁边生成一个config.cfg  
随便使用一个文本编辑软件(比如记事本)打开它  
找到path 填入 C:/Users/你的Windows用户名/AppData/Local/Oculus/DeviceCache.json 保存后再次启动main.exe  
Web端会出现在 http://127.0.0.1:23333/2SensorStatus   

### 如何使用/How to use
#### 浏览器直接打开
> 随便找个浏览器直接打开 http://127.0.0.1:23333/2SensorStatus 就能看到
#### OBS中调用
> 打开OBS，添加 来源-浏览器  
URL处填http://127.0.0.1:23333/2SensorStatus  
宽度240 高度130 并删除自定义CSS框中的全部内容  
勾选不可见时关闭源后点击确定即可导入  
