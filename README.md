
# Oculus Rift CV1 传感器追踪状态检查工具  
# (Oculus Rift CV1 Sensor Tracking Status Checking Tool)  
这是一个可以显示单独的Oculus传感器追踪状态的工具，使用web显示，可以轻松地导入OBS  
This is a tool that can show you individual oculus sensor tracking status, output as a web, easy cooperate with OBS

### 初始化/initialization
> [Releases](https://github.com/VAIO-Dong/Oculus-Rift-CV1-Sensor-Tracking-Status-Checking-Tool/releases)  
你可以从上面的Releases中下载到我们打包好的整个程序  
找个你喜欢的地方建个文件夹把所有东西解压进去 双击main.exe启动  
输入传感器1的序列号然后回车（序列号可以在Oculus的应用程序里看）  
输入传感器2的序列号然后回车  
输入EXIT回车（3/4传感器的web端还没做XD） 
然后找到在他的旁边生成一个config.cfg  
理论上这个时候程序会自动获取到你的Oculus日志文件并填入到config.cfg中  
(Oculus日志文件一般在‘C:/Users/你的Windows用户名/AppData/Local/Oculus/DeviceCache.json’)  
  ---如果需要修改的话可以随便使用一个文本编辑软件(比如记事本)打开config.cfg  
  ---path代表着Oculus日志文件的位置 url是请求生成的web的位置 host是监听的主机 device_id是传感器的序列号  
如果能正常识别自动配置而且默认配置都未更改
> Web端会出现在 http://127.0.0.1:23333/SensorStatus   

### 如何使用/How to use
#### 浏览器直接打开
> 在运行程序的主机上随便找个浏览器直接打开 http://127.0.0.1:23333/SensorStatus 就能看到
#### OBS中调用
> 打开OBS，添加 来源-浏览器  
URL处填 http://127.0.0.1:23333/SensorStatus  
宽度240 高度130 并删除自定义CSS框中的全部内容  
勾选不可见时关闭源后点击确定即可导入  
#### 手机监看
> 用记事本打开config.cfg 找到"host":"127.0.0.1:23333"字样 更改为"0.0.0.0:23333"然后保存  
重启main.exe 如果这个时候有防火墙弹出请允许main.exe访问网络  
在同个局域网下的手机直接用浏览器访问 http://运行程序的那台电脑的IP地址:23333/SensorStatus 即可  
#### 双机直播时OSB调用
> 程序的设置步骤与手机监看相同 执行完后到另外一台电脑上的打开OBS 添加 来源-浏览器  
URL处填 http://运行程序的那台电脑的IP地址:23333/SensorStatus  
宽度240 高度130 并删除自定义CSS框中的全部内容  
勾选不可见时关闭源后点击确定即可导入  
