//index.js
//获取应用实例
const app = getApp()
const fileManager = wx.getFileSystemManager();
Page({
  data: {
    motto: '笔画练字',
    userInfo: {},
    hasUserInfo: false,
    canIUse: wx.canIUse('button.open-type.getUserInfo'),
    pageData: {},
    hanzi:'',
    search:''
  },
  onShow() {
    if(app.globalData.hanzi) {
      this.getHanziData();
      app.globalData.hanzi = ""
    }
    if (typeof this.getTabBar === 'function' &&
      this.getTabBar()) {
      this.getTabBar().setData({
        selected: 0
      })
    }
  },
  click_chengyu(e){
    app.globalData.hanzi = e.currentTarget.dataset.zi
    wx.switchTab({
      url: '/pages/chengyu/index'
    })
  },
  onLoad: function (e) {
    wx.showShareMenu({

      withShareTicket:true,
      
      menus:['shareAppMessage','shareTimeline']
      
    });
    this.getHanziData();
    if (app.globalData.userInfo) {
      this.setData({
        userInfo: app.globalData.userInfo,
        hasUserInfo: true
      })
    } else if (this.data.canIUse){
      // 由于 getUserInfo 是网络请求，可能会在 Page.onLoad 之后才返回
      // 所以此处加入 callback 以防止这种情况
      app.userInfoReadyCallback = res => {
        this.setData({
          userInfo: res.userInfo,
          hasUserInfo: true
        })
      }
    } else {
      // 在没有 open-type=getUserInfo 版本的兼容处理
      wx.getUserInfo({
        success: res => {
          app.globalData.userInfo = res.userInfo
          this.setData({
            userInfo: res.userInfo,
            hasUserInfo: true
          })
        }
      })
    }
  },
  getUserInfo: function(e) {
    console.log(e)
    app.globalData.userInfo = e.detail.userInfo
    this.setData({
      userInfo: e.detail.userInfo,
      hasUserInfo: true
    })
  },

  getHanziData: function(e) {
    self = this;
    var search = self.data.search;
    if(app.globalData.hanzi) {
      search = app.globalData.hanzi
    }
    wx.request({
      url: '',
      data: {
        token:''
      },
      success: function(res) {
        var mp3 = res.data.mp3;
        var pinyin = res.data.pinyin;
        var pinyins = pinyin.split(",");
        var audios = mp3.split(",");
        var chengyus = res.data.chengyu.split(",");
        res.data.chengyus = chengyus;
        var data = [];
        for(var index = 0; index < pinyins.length; index++) {
          if(pinyins[index] == "") {
            break;
          }
          var tmp = {pinyin:pinyins[index], audio:audios[index]};
          data[index] = tmp;
        }
        res.data.pinyins = data;
        self.setData({pageData:res.data});
        //初始化这个玩意
        //app.globalData.hanzi = '';
      },
      error: function(res) {
      }
    })
  },
  searchHanzi: function(e) {
    var hanzi = this.data.hanzi;
    if (hanzi == "") {
      return;
    }
    this.setData({search:hanzi})
    this.getHanziData();
  },
  inputHanzi: function(e) {
    this.setData({hanzi:e.detail.value})
  }
})
