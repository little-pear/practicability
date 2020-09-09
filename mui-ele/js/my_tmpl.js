(function(){
	window.my_tmpl = {
		//创建item 模版
		itemTmpl:function(obj){
			
			if (obj.isDistribution == 0) {
				var distribution = '免配送费'
			} else {
				var distribution = obj.isDistribution + '元配送费'
			}
			
			//计算分钟
			// 1分钟=50米
			var i = 50
			var m = obj.distance/i
			
			var tmpl = '<div class="wm-list-div" id_data="'+obj.id+'">'
						+'	<div class="wm-list-div-l">'
						+'		<img src="images/logo.png" >'
						+'	</div>'
						+'	<div class="wm-list-div-r">'
						+'		<div class="wm-list-div-r-o">'
						+'			<div class="wm-list-div-r-o-l">'
						+'				<h4>'+obj.title+'</h4>'
						+'			</div>'
						+'			<div class="wm-list-div-r-o-r">'
						+'				<h5> <span>¥</span><span class="big-txt">'+obj.price+'</span>元起送</h5>'
						+'			</div>'
						+'		</div>'
						+'		<div class="wm-list-div-r-t">'
						+'			<div class="wm-list-div-r-t-l">'
						+'				<span class="mui-icon mui-icon-star-filled start-ys start-active"></span>'
						+'				<span class="mui-icon mui-icon-star-filled start-ys start-active"></span>'
						+'				<span class="mui-icon mui-icon-star-filled start-ys start-active"></span>'
						+'				<span class="mui-icon mui-icon-star-filled start-ys start-active"></span>'
						+'				<span class="mui-icon mui-icon-star-filled start-ys"></span>'
						+'				<span class="small-txt o-txt">4.5</span>'
						+'				<span class="small-txt">月售'+obj.buyNumber+'件</span>'
						+'			</div>'
						+'			<div class="wm-list-div-r-t-r">'
						+'				<span class="small-txt">'+distribution+'</span>'
						+'			</div>'
						+'		</div>'
						+'		<div class="wm-list-div-r-tt">'
						+'			<span  class="small-txt">'+m+'分钟</span> <span  class="small-txt">'+obj.distance+'米</span>'
						+'		</div>'
						+'	</div>'
						+'</div>';
			return tmpl;
		}
	}
	
})();