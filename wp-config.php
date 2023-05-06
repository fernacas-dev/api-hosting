<?php
define('DB_NAME', 'wordpress488530');
define('DB_USER', 'root');
define('DB_PASSWORD', 'DontTouchMyDbServer2021*');
define('DB_HOST', '172.17.0.8');
define('DB_CHARSET', 'utf8mb4');
define('DB_COLLATE', '');

define('AUTH_KEY',         '0Io1!G{`|<b*lSQK-po%QUlDKv8qC?j3?dyQ70>?ChHWSHDccc=7hioHG24~<fvb');
define('SECURE_AUTH_KEY',  'pFj0Qq~Do*@Fr90(j.IJ&voKJ3nHiZ,m?wF E^*/Y>,*`k6x/Qe#@2uHwaVb.Fji');
define('LOGGED_IN_KEY',    'sd.<<uoG.unk?QxZ_XuK_K+D|FBLX;NXm>`Q*AI#~t/#d342:dE/(/KUput$Xz8O');
define('NONCE_KEY',        '<YJ&pg*KFHAxVg8i=nM|P$w_HwK/1,A]>/ls>n}(FUm$yQ$0FTC@*h-}Kl5%FJ@t');
define('AUTH_SALT',        '!a[m=zC>F9JEO+>Dg?h%Zp!6}Y&<30un1~c7tQ~47m6-yv!$BtBkix(1$?Y7?+zJ');
define('SECURE_AUTH_SALT', '1pt.@U|]Ji!%71$dM1Zdx;4O%)2}baWo6_`i9f=<P:G,)_K2+<5rlG,UWc~]##76');
define('LOGGED_IN_SALT',   'Z,W`&w%d3.*F,{d!+3$Ru`3kiP A,#K9mgVquC&J1c/fa}4I};DQ(V1MqsVK3f#Y');
define('NONCE_SALT',       'wUa^^DBq[87OJs*dyN@w!]Q|f4Xu8Dko_)V~Xlw(x 6I7`mjM[JZ5-zY4M[s}F k');

$table_prefix = 'wp_';

define('WP_DEBUG', false);

if (!defined('ABSPATH')) {
	define('ABSPATH', __DIR__ . '/');
}

require_once ABSPATH . 'wp-settings.php';
