//获取操作系统类型
export function GetSystemType(){
  var sUserAgent = navigator.userAgent;
  var isWin = (navigator.platform == "Win32") || (navigator.platform == "Windows");
  var isMac = (navigator.platform == "Mac68K") || (navigator.platform == "MacPPC") || (navigator.platform == "Macintosh") || (navigator.platform == "MacIntel");
  if (isMac) return "macos";
  var isUnix = (navigator.platform == "X11") && !isWin && !isMac;
  if (isUnix) return "unix";
  var isLinux = (String(navigator.platform).indexOf("Linux") > -1);
  if (isLinux) return "linux";
  if (isWin) {
      return "windows"
  }
  return "other";
}
