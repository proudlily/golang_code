jQuery(function($) {
	$(document).ready(function() {
		/* SETUP AND DEBUG */
		if ($('#wpadminbar').length > 0) { $('body').addClass('wpadmin-correct') }; // Reposition in admin mode to account for WPAdminBar.
		
		$('.archive .casestudy-selector .panel:first').addClass('active');


		/* COLUMIZE FEATURED NEWS ARTICLE */
		$('.autocolumn').columnize({
			columns: 2,
			lastNeverTallest: true
		});


		/* MOBILE MENU  */ 
		$('.mobile-menu-trigger a.menu').click(function(event){
			event.preventDefault();
            $('.primary-nav > nav ul').slideToggle();
		});
		$(window).on('resize', function(){
			var win = $(this);
      		if (win.width() >= 830) { $('.primary-nav > nav ul').show(); }
      		else { $('.primary-nav > nav ul').hide(); };
		});


		/* MEDIA ELEMENT */
		if ( $('body').hasClass('home') ){
			$('body.home .panel.video video').mediaelementplayer({
			    enablePluginSmoothing: true,
	        	enableAutosize: false,
			    features: ['playpause','progress','current','duration','fullscreen'],
			    alwaysShowControls: false
			});
		};


		/* MAILING LIST FORM FIELDS */	
  		$('label').inFieldLabels();


		/* LARGE FORMAT CONTENT SLIDER */		
	    var slideContainer = $('.large-format-slider .slider');
	    var sliderControls = true;
		slideContainer.imagesLoaded(function(){
	        slideContainer.carouFredSel({
	            width       :   '100%',
	            responsive  :   true,
	            prev		: 	{ button			: function() { return $(this).closest('.large-format-slider').find('.prev'); }},
		    	next		: 	{ button			: function() { return $(this).closest('.large-format-slider').find('.next'); }},
	            scroll      :   {   pauseOnHover    : true,
	            					easing			: 'easeInOutQuart',
	            					pauseOnHover 	: true,
	            					duration 		: 750,
	            					onBefore 		: function() { 
	            										var slideNum = slideContainer.triggerHandler( 'currentPosition' );
														$('.archive .casestudy-selector .panel').removeClass('active');
														$('.archive .casestudy-selector .panel').eq(slideNum).addClass('active');
	            					}},
	            auto        :   {   play            : true,
	                                timeoutDuration : 5000 },
	            swipe       :   {   onTouch         : false,
	                                onMouse         : false },
	            pagination 	: 	'.slide-pager',
	            onCreate: function () {
                $(window).on('resize', function () {
                  slideContainer.parent().add(slideContainer).height(slideContainer.children().first().height());
                }).trigger('resize');
              }
	        });
			slideContainer.swipe({
				excludedElements : "button, input, select, textarea, .noSwipe",
				swipeLeft: function() {
					slideContainer.trigger('next', 1);
				},
				swipeRight: function() {
					slideContainer.trigger('prev', 1);
				},
				tap: function(event, target) {
					event.preventDefault();
					var $target = $(target);
					if($target.is('a[href="#"]')) {
						$target.trigger('click');
					} else {
						window.open($target.closest('.slide').find('a').attr('href'), '_self');
					}
				}
			});
		});


		/* PARALLAX IMAGES */
		if( jQuery.browser.mobile ){
			$('.single-newsarticle .featured .page-content .parallax-window').css( {
				'background-attachment': 'scroll',
				'padding-bottom': '80%',
				'background-position': 'center top',
				'height': '0',
				'background-size': 'cover'
			});
		}
		else {
			$.stellar({
				horizontalScrolling: false,
				verticalScrolling: true,
				horizontalOffset: 0,
				verticalOffset: 0,
				responsive: true,
				scrollProperty: 'scroll',
				positionProperty: 'position',
				parallaxBackgrounds: true,
				parallaxElements: false,
				hideDistantElements: false,
				hideElement: function($elem) { $elem.hide(); },
				showElement: function($elem) { $elem.show(); }
			});
		}
	    

	    /* ISOTOPE LAYOUT */
	    var $isotopeContainer = $('.isotope');
	    var $window = $(window);
		$isotopeContainer.imagesLoaded(function(){
			// $isotopeContainer.packery({
			// 	itemSelector 	: '.panel',
			// 	columnWidth: 204,
			// 	isResizeBound: true
			// });
			if($window.width()<830){
			  	$isotopeContainer.isotope({
		    		// update columnWidth to a percentage of container width
    				masonry: { columnWidth: $isotopeContainer.width() / 3 }
  				});
			}
			else {
				$isotopeContainer.isotope({
		    		// update columnWidth to a percentage of container width
    				masonry: { columnWidth: 408 }
  				});
			}
			
			$window.smartresize(function(){
				if($window.width()<830){
				  	$isotopeContainer.isotope({
			    		// update columnWidth to a percentage of container width
	    				masonry: { columnWidth: $isotopeContainer.width()}
	  				});
				}
				else {
					$isotopeContainer.isotope({
			    		// update columnWidth to a percentage of container width
	    				masonry: { columnWidth: 408 }
	  				});
				}
			});
		});

		$(document).on('change', '.isotope-filter', function(event) {
			event.preventDefault();
			var val = $(this).val();

			if(val === 'all' ) {
				$isotopeContainer.isotope({
					filter: '[data-categories]'
				});
			}
			else {
				$isotopeContainer.isotope({
					filter: '[data-categories*='+val+']'
				});
			}
		});

		var $body = $('body'), $html = $('html');

		if($body.hasClass('single-casestudy') && !$html.hasClass('csscalc')) {
			var elem = $('.casestudy-summary .inner-wrapper');
			$window.resize(function(event) {
				if($window.width() < 830) {
					elem.css('width', (elem.closest('.container').width() - 102) + 'px');
				}
			});

			if($window.width() < 830) {
				elem.css('width', (elem.closest('.container').width() - 102) + 'px');
			}
		}

		if($body.hasClass('single-casestudy') && !$html.hasClass('csscalc')) {
			var elem = $('.sidebar .search form div input');
			$window.resize(function(event) {
				if($window.width() < 830) {
					elem.css('width', (elem.parent().width() - 102) + 'px');
				}
			});

			if($window.width() < 830) {
				elem.css('width', (elem.parent().width() - 102) + 'px');
			}
		}
	});	
});