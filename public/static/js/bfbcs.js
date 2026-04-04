/* © Copyright by Dominik "1ApRiL" Herbst */

var bfbcs = {
	configure: function() {
		$.ajaxSetup({
			url: location.pathname,
			async: true,
			dataType: 'json',
			type: 'POST'
		});
		
	},
	qtout:false,
	lastpos:-1,
	addPlayerQueue: function(obj) {
		if($.browser.msie && parseInt($.browser.version)<7) {
			alert('Your Browser is too old. Please update your internet browser or use another one!');
			return;
		}
		
		var obj = $(obj);
		obj.addClass('loading');
		var refreshtime=10000;
		
		var func = function(dat) {
			obj.removeClass('loading');
			if(dat && dat.timediff) {
				if(bfbcs.lastpos!=-1) {
					if(bfbcs.qtout) { window.clearTimeout(bfbcs.qtout); bfbcs.qtout=false; }
					bfbcs.qtout = window.setTimeout(function(){ window.location.reload(); },refreshtime);
				}
				var d = obj.html();
				var clt = bfbcs.queuewait;
				d=d.replace(/\(.+?\)/,'('+clt+')');
				obj.html(d);
			}
			if(dat && dat.added && dat.count) {
				var d = obj.html();
				var clt = bfbcs.queuebtntext;
				if(dat.timeleft) {
					clt=bfbcs.queuebtntexttleft;
					clt=clt.replace('TIMEL',dat.timeleft);
				}
				clt=clt.replace('POS',dat.added);
				clt=clt.replace('COUNT',dat.count);
				d=d.replace(/\(.+?\)/,'('+clt+')');
				obj.html(d);
				bfbcs.lastpos=dat.added;
				if(bfbcs.qtout) { window.clearTimeout(bfbcs.qtout); bfbcs.qtout=false; }
				bfbcs.qtout = window.setTimeout(function(){ bfbcs.qtout=false; bfbcs.addPlayerQueue(obj); },refreshtime);
			} else if(dat && dat.added===false && dat.count) {
				var d = obj.html();
				var clt = bfbcs.queuefulltext;
				clt=clt.replace('COUNT',dat.count);
				d=d.replace(/\(.+?\)/,'('+clt+')');
				obj.html(d);
				if(bfbcs.qtout) { window.clearTimeout(bfbcs.qtout); bfbcs.qtout=false; }
				bfbcs.qtout = window.setTimeout(function(){ bfbcs.qtout=false; bfbcs.addPlayerQueue(obj); },refreshtime);
			}
		};
		var err = function(reqs) {
			obj.removeClass('loading');
			if(reqs.status && reqs.status==500) {
				if(bfbcs.qtout) { window.clearTimeout(bfbcs.qtout); bfbcs.qtout=false; }
				bfbcs.qtout = window.setTimeout(function(){ bfbcs.qtout=false; bfbcs.addPlayerQueue(obj); },refreshtime);
			}
		};
		
		$.ajax({
			data:{request:'addplayerqueue',pcode:bfbcs.pcode},
			success:func,
			error:err
		});
		
	},
	
	currDetail:{},
	currDetailActive:{},
	showDetail: function(typ,id,idactive,acobj) {
		var func = function() {
			var o = $('#'+id);
			var pr=o.parents('.stcol_2');
			var ph=o;
			var oa = $('#'+idactive);
			bfbcs.currDetail[typ]=o;
			bfbcs.currDetailActive[typ]=oa;
			o.show();
			
			
			ph.css({'margin-top':0});
			
			var pp=pr.offset();
			var op=o.offset();
			var oap=oa.offset();
			var oapo=oa.parents('.stcol_1').first();
			var oapop=oapo.offset();
			
			var st=$(window).scrollTop();
			var wh=$(window).height();
			var mint=st+10;
			var maxt=st+wh-10;
			var oh=ph.outerHeight();
			//alert(oh+' = '+o.parent().outerHeight()+'+'+ph.outerHeight());
			
			var mt=Math.floor(oap.top-oh/2-pp.top);
			
			if(mt<0) mt=0;
			else if(pp.top+mt<mint) mt=mint-pp.top;
			
			if(pp.top+mt+oh>maxt) mt-=pp.top+mt+oh-maxt;
			if(mt<0) mt=0;
			
			if(pp.top+mt+oh>oapop.top+oapo.height()) mt-=pp.top+mt+oh-(oapop.top+oapo.height());
			if(mt<0) mt=0;
			
			ph.css({'margin-top':mt+'px'});
			
			oa.addClass('active');
		};
		
		if(typeof(bfbcs.currDetail[typ])!='undefined') {
			bfbcs.currDetail[typ].css({'margin-top':0});
			bfbcs.currDetail[typ].hide();
		} else {
			$('#'+typ+'_').hide();
		}
		if(typeof(bfbcs.currDetailActive[typ])!='undefined') {
			bfbcs.currDetailActive[typ].removeClass('active');
		}
		func();
		
		if(typeof(acobj)!='undefined') {
			var to=$(acobj);
			var o = $('#'+id).clone();
			$('#content').append(o);
			
			o.addClass('movable cont');
			
			
			var posobj=function(x,y) {
				var topos={};
				if(typeof(x)=='undefined') {
					topos=to.offset();
					topos.left+=1;
					topos.top+=to.height()+1;
				} else {
					topos.left=x;
					topos.top=y;
				}
				
				topos.left+=10;
				topos.top+=10;
				
				var mt=$(window).scrollTop()+$(window).height()-5;
				var ml=$(window).scrollLeft()+$(window).width()-5;
				
				if(topos.top+o.height()>mt) {
					topos.top-=topos.top+o.height()-mt;
				}
				if(topos.left+o.width()>ml) {
					topos.left-=topos.left+o.width()-ml;
				}
				o.offset(topos);
			};
			posobj();
			
			to.mousemove(function(e){ posobj(e.pageX,e.pageY); });
			to.mouseout(function(){ o.remove(); });
		}
		
	},
	
	addFavorite: function(name) {
		var succf = function(dat) {
			if(dat.code=='failed') return;
			var nel = $('<a>');
			nel.attr('href',dat.url);
			nel.addClass('lvl1');
			nel.attr('pname',name);
			nel.html(name);
			
			$('.favplayers').append(nel);
			$('#favplayeradd').hide();
			$('#favplayerdel').show();
			$('#favplayeradd').removeClass('loading');
			
		};
		var errf = function() {
			$('#favplayeradd').removeClass('loading');
		};
		
		$('#favplayeradd').addClass('loading');
		
		$.ajax({
			data:{globalrequest:'addfavorite',name:name},
			success:succf,
			error:errf
		});
	},
	delFavorite: function(name) {
		var succf = function(dat) {
			if(dat.code=='failed') return;
			var objs = $('.favplayers a');
			objs.each(function(){
				if($(this).attr('pname') && $(this).attr('pname')==name) {
					$(this).remove();
					$('#favplayerdel').hide();
					$('#favplayeradd').show();
					$('#favplayerdel').removeClass('loading');
				}
			});
			
		};
		var errf = function() {
			$('#favplayerdel').removeClass('loading');
		};
		
		$('#favplayerdel').addClass('loading');
		
		$.ajax({
			data:{globalrequest:'delfavorite',name:name},
			success:succf,
			error:errf
		});
	},
	loadFavourite: function() {
		var dela=$('<a href="#" class="delbtn"></a>');
		dela.css({'display':'block','height':'16px','position':'absolute'});
		dela.hide();
		$(document.body).append(dela);
		
		var objs = $('.favplayers a');
		objs.each(function(){
			var po=$(this);
			if(po.attr('pname')) {
				var coords=po.offset();
				var mover=function(){
					var ncoords=po.offset();
					ncoords.left+=po.outerWidth()-dela.outerWidth();
					ncoords.top+=1;
					dela.show();
					dela.offset(ncoords);
					
					//alert(ncoords.left+'+='+po.outerWidth()+'-'+dela.outerWidth()+'\n'+ncoords.top);
					
					dela.unbind('click');
					dela.click(function(){
						bfbcs.delFavorite(po.attr('pname'));
					});
				}
				po.mouseover(mover);
				po.mouseout(function(e){
					if((e.pageY<coords.top || e.pageY>coords.top+dela.outerHeight()) && (e.pageX<coords.left || e.pageX>coords.left+dela.outerWidth())) {
						dela.hide();
						dela.unbind('click');
					}
				});
				
			}
		});
	},
	grsite: function(objid,letid,start) {
		var obj = $(objid);
		var objlet = $(letid);
		objlet.addClass('loading');
		
		obj.animate({opacity:0.0},500);
		
		var filter = obj.attr('grfilter');
		if(filter.indexOf('{')==0) {
			filter = JSON.parse(filter);
		}
		
		$.ajax({
		data:{request:obj.attr('grequest'), filter:filter, start: start},
		success: function(res) { 
			objlet.removeClass('loading');
			if(res && res.content) {
				objlet.html('');
				
				obj.html(res.content);
				obj.animate({opacity:1.0},500);
				
				if(res.letters.back) {
					objlet.append('<a href="javascript:bfbcs.grsite(\''+objid+'\',\''+letid+'\','+res.letters.back_start+')" class="back">'+res.text_back+'</a>');
				}
				if(res.letters.start_page) {
					objlet.append('<a href="javascript:bfbcs.grsite(\''+objid+'\',\''+letid+'\','+res.letters.start_start+')" class="back">'+res.letters.start_page+'</a> .. ');
				}
				
				if(res.letters.pagecount>1) {
					for(var site in res.letters.pages) {
						objlet.append('<a href="javascript:bfbcs.grsite(\''+objid+'\',\''+letid+'\','+res.letters.pages[site]+')" class="let'
							+(res.letters.curr==site?' active':'')+'">'+site+'</a>'
						);
					}
				}
				
				if(res.letters.end_page) {
					objlet.append(' .. <a href="javascript:bfbcs.grsite(\''+objid+'\',\''+letid+'\','+res.letters.end_start+')" class="back">'+res.letters.end_page+'</a>');
				}
				if(res.letters.next) {
					objlet.append('<a href="javascript:bfbcs.grsite(\''+objid+'\',\''+letid+'\','+res.letters.next_start+')" class="next">'+res.text_next+'</a>');
				}
				
			}
		}
	});
		
		
	}
	
};

bfbcs.getGraphic = function(id) {
	var dialog = $('#getgraphicdialog');
	
	dialog.dialog({
		modal:true,
		width: 600,
		maxWidth:1000
	});
	dialog.dialog('open');
	
	dialog.css({opacity:0.1});
	
	$.ajax({
		data:{request:'getgraphic', id: id},
		success: function(res) { 
			dialog.animate({opacity:1.0},500);
			$('.title',dialog).html(res.title);
			$('.author',dialog).html(res.author);
			$('#previewgraphic').attr('src',res.img);
			$('#previewgraphic').animate({'width':res.width,'height':res.height},400);
			$('.width',dialog).html(res.width);
			$('.height',dialog).html(res.height);
			$('.likes',dialog).html(res.likes);
			$('.dislikes',dialog).html(res.dislikes);
			$('.rating',dialog).html(res.rating);
			
			$('.bbcode',dialog).val(res.bbcode);
			$('.htmlcode',dialog).val(res.html);
			$('.url',dialog).val(res.img);
			
			var obj = $('.playername',dialog);
			obj.html(res.player.name);
			obj.addClass('playername');
			obj.css('background-image','url('+res.player.rankimg+')');
			
			dialog.attr('graphicid',id);
			
		}
	});
};

bfbcs.choosePlayer = function(obj,plat) {
	obj = $(obj);
	
	var dialog = $('#chooseplayer_dialog').clone().appendTo('body');
	var playfield = $('input:text[name=\'player\']',dialog);
	var platfield = $('select[name=\'platform\']',dialog);
	$('select[name=\'platform\'] option:selected',dialog).attr('selected',false);
	$('select[name=\'platform\'] option[value=\''+plat+'\']',dialog).attr('selected',true);
	
	if(obj.hasClass('playername')) {
		playfield.val(obj.text());
	}
	
	if(bfbcs.editor && bfbcs.editor.id) {
		var edid = bfbcs.editor.id;
	} else {
		var edid = $('#getgraphicdialog').attr('graphicid');
	}
	
	dialog.dialog({ 
		modal: true,
		buttons: { 
			"Ok": function() {
				$.ajax({
					data:{request:'setplayer', name: playfield.val(), platform: platfield.val(), id:edid},
					success: function(res) { 
						if(res && res.player) {
							obj.html(res.player.name);
							obj.addClass('playername');
							obj.css('background-image','url('+res.player.rankimg+')');
							
							$('#previewgraphic').attr('src',res.player.img)
							
							dialog.dialog("close");
						}
					},
					error: function() { 
						dialog.dialog("close");
					}
				});
				
			},
			"Cancel": function() { $(this).dialog("close"); } 
		} 
	});
	
};
bfbcs.rateGrpahic = function(obj,like) {
	obj = $(obj);
	var edid = $('#getgraphicdialog').attr('graphicid');
	obj.addClass('loading');
	$.ajax({
		data:{request:'rategrpahic', id:edid, like: like},
		success: function(res) {
			obj.removeClass('loading');
			if(res.successful) {
				obj.addClass('ok');
				window.setTimeout(function() { obj.removeClass('ok'); },1000);
			}
		}
	});


};
bfbcs.autoDateTime = function(objs) {
	
	objs.each(function(oi,el) {
		var o = $(el);
		var val = o.html();
		var format = 0;
		var dt = new Date();
		var offset = dt.getTimezoneOffset();
		if(o.attr('pasttime')) {
			dt.setTime(o.attr('pasttime')*1000);
		} 
		if(val.match(/^(\d\d\d\d)-(\d\d)-(\d\d) (\d\d):(\d\d)/)) {
			format=1;
			dt=new Date(RegExp.$1+'-'+RegExp.$2+'-'+RegExp.$3+'T'+RegExp.$4+':'+RegExp.$5+':00Z');
			
		}
		
		if(!format) return;
		/*
		if(offset) {
			var csec = dt.getTime();
			var oldsec = csec;
			csec += offset*60*1000*-1;
			dt.setTime(csec);
		}*/
		
		if(o.hasClass('gerfmt')) {
			format=2;
		}
		
		var ret='';
		if(format==1) {
			ret += dt.getFullYear()+'-';
			ret += (dt.getMonth()+1<10?'0'+(dt.getMonth()+1):dt.getMonth()+1)+'-';
			ret += (dt.getDate()<10?'0'+dt.getDate():dt.getDate())+' ';
			ret += (dt.getHours()<10?'0'+dt.getHours():dt.getHours())+':';
			ret += (dt.getMinutes()<10?'0'+dt.getMinutes():dt.getMinutes());
		}
		if(format==2) {
			var wnames = ["So","Mo","Di","Mi","Do","Fr","Sa"];
			ret += wnames[dt.getDay()]+' ';
			ret += (dt.getDate()<10?'0'+dt.getDate():dt.getDate())+'. ';
			var mnames = ["Jan","Feb","Mär","Apr", "Mai", "Jun","Jul", "Aug", "Sep","Okt", "Nov", "Dez"];
			ret += mnames[dt.getMonth()]+' ';
			ret += dt.getFullYear()+' ';
			ret += (dt.getHours()<10?'0'+dt.getHours():dt.getHours())+':';
			ret += (dt.getMinutes()<10?'0'+dt.getMinutes():dt.getMinutes());
		}
		//alert(val+'\n'+ret);
		o.html(ret);
		
		if(o.attr('pasttime')) {
			var nd=new Date();
			var diff=Math.floor((nd.getTime()-dt.getTime())/1000);
			if(diff<60) {
				o.html(diff+' sec ago');
			} else if(diff<3600) {
				var i=Math.floor(diff/60);
				var sec=diff-i*60;
				o.html(Math.floor(diff/60)+'m '+sec+'s ago');
			}
		}
		
	});
	
};

bfbcs.initHover = function(obj) {
	
	obj.bind('mouseover',function(e) {
		//
		var offs = $(this).offset();
		var hoverid = $(this).attr('bhover');
		if(hoverid) {
			var o = $(hoverid);
			if(o) {
				var newo = o.clone();
				newo.appendTo('body');
				newo.addClass('hoverbox');
				newo.attr('id','');
				newo.css('display','block');
				
				
				var mmove = function(e) {
					var nx = e.pageX+10;
					var ny = e.pageY+10;
					var maxx = $(window).width()+$(window).scrollLeft()-1;
					var maxy = $(window).height()+$(window).scrollTop()-1;
					if(nx+newo.outerWidth()>maxx)
						nx = maxx-newo.outerWidth();
					else if(ny+newo.outerHeight()>maxy)
						ny = maxy-newo.outerHeight();
					
					newo.css('left',nx);
					newo.css('top',ny);
				};
				
				mmove(e);
				
				var mout = function() {
					newo.remove();
					$(this).unbind('mousemove',mmove);
					$(this).unbind('mouseout',mout);
				};
				
				$(this).bind('mousemove',mmove);
				$(this).bind('mouseout',mout);
				
			}
		}
		
	});
	
};

bfbcs.count = function(name) {
	$.ajax({
		data:{globalrequest:'countit',name:name}
	});
};

bfbcs.configure();


var initPStatsRow=function(){
	var rowobj=$('#pstatsrow');
	
	var initmenu=function(mo,dropo){
		if(dropo.size()==0) return;
		var tout=null;
		var cltoutfunc=function(){
			if(tout) {
				clearTimeout(tout);
				tout=null;
			}
		};
		
		var arrow=$('<div></div>');
		mo.append(arrow);
		
		var pos=mo.position();
		pos.top+=mo.outerHeight();
		dropo.css({position:'absolute', left:pos.left+'px', top:pos.top+'px', 'min-width':(mo.outerWidth()-8)+'px'});
		
		mo.mouseover(function(){
			cltoutfunc();
			dropo.show();
			mo.addClass('active');
			var hidefunc=function(){
				tout=setTimeout(function(){
					dropo.hide();
					mo.removeClass('active');
				},50);
			};
			
			dropo.one('mouseover',cltoutfunc);
			dropo.one('mouseleave',hidefunc);
			mo.one('mouseleave',hidefunc);
		});
		
		
		
	};
	
	$('a',rowobj).each(function(i){
		var o=$(this);
		var classes=o.attr('class');
		if(classes) classes=classes.split(' ');
		else classes=[];
		for(var i in classes) {
			if(classes[i] && classes[i].length)  initmenu(o,$('#'+classes[i])); 
		}
		
	});
	
};

$(document).ready(function(){
	initPStatsRow();
	if(typeof($.tablesorter)!='undefined') {
		$.tablesorter.addParser({ 
			id: 'dps', 
			is: function(s) {
				if(s.match(/^\d+( \d+)+$/) || s.match(/^\d+h \d+m \d+s$/)) {
					return true; 
				}
				return false;
			}, 
			format: function(s) { 
				if(s.match(/^\d+h \d+m \d+s$/)) {
					s=s.replace(/h /g,'.'); 
					s=s.replace(/m /g,''); 
					s=s.replace(/s/g,''); 
				}
				return s.replace(/ /g,'')*1; 
			}, 
			type: 'numeric' 
		}); 
	}
	
	var objs = $('.sortable');
	if(objs.tablesorter) objs.tablesorter();
	
	bfbcs.initHover($('div[bhover]'));
	bfbcs.autoDateTime($('.datetime'));
	
	bfbcs.loadFavourite();
	
});
