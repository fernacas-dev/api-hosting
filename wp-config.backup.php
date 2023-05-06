<?php
/**
 * The base configuration for WordPress
 *
 * The wp-config.php creation script uses this file during the installation.
 * You don't have to use the web site, you can copy this file to "wp-config.php"
 * and fill in the values.
 *
 * This file contains the following configurations:
 *
 * * MySQL settings
 * * Secret keys
 * * Database table prefix
 * * ABSPATH
 *
 * @link https://wordpress.org/support/article/editing-wp-config-php/
 *
 * @package WordPress
 */

// ** MySQL settings - You can get this info from your web host ** //
/** The name of the database for WordPress */
define( 'DB_NAME', 'wordpress488530' );

/** MySQL database username */
define( 'DB_USER', 'root' );

/** MySQL database password */
define( 'DB_PASSWORD', 'DontTouchMyDbServer2021*' );

/** MySQL hostname */
define( 'DB_HOST', '172.17.0.8' );

/** Database charset to use in creating database tables. */
define( 'DB_CHARSET', 'utf8mb4' );

/** The database collate type. Don't change this if in doubt. */
define( 'DB_COLLATE', '' );

/**#@+
 * Authentication unique keys and salts.
 *
 * Change these to different unique phrases! You can generate these using
 * the {@link https://api.wordpress.org/secret-key/1.1/salt/ WordPress.org secret-key service}.
 *
 * You can change these at any point in time to invalidate all existing cookies.
 * This will force all users to have to log in again.
 *
 * @since 2.6.0
 */
define( 'AUTH_KEY',         '0Io1!G{`|<b*lSQK-po%QUlDKv8qC?j3?dyQ70>?ChHWSHDccc=7hioHG24~<fvb' );
define( 'SECURE_AUTH_KEY',  'pFj0Qq~Do*@Fr90(j.IJ&voKJ3nHiZ,m?wF E^*/Y>,*`k6x/Qe#@2uHwaVb.Fji' );
define( 'LOGGED_IN_KEY',    'sd.<<uoG.unk?QxZ_XuK_K+D|FBLX;NXm>`Q*AI#~t/#d342:dE/(/KUput$Xz8O' );
define( 'NONCE_KEY',        '<YJ&pg*KFHAxVg8i=nM|P$w_HwK/1,A]>/ls>n}(FUm$yQ$0FTC@*h-}Kl5%FJ@t' );
define( 'AUTH_SALT',        '!a[m=zC>F9JEO+>Dg?h%Zp!6}Y&<30un1~c7tQ~47m6-yv!$BtBkix(1$?Y7?+zJ' );
define( 'SECURE_AUTH_SALT', '1pt.@U|]Ji!%71$dM1Zdx;4O%)2}baWo6_`i9f=<P:G,)_K2+<5rlG,UWc~]##76' );
define( 'LOGGED_IN_SALT',   'Z,W`&w%d3.*F,{d!+3$Ru`3kiP A,#K9mgVquC&J1c/fa}4I};DQ(V1MqsVK3f#Y' );
define( 'NONCE_SALT',       'wUa^^DBq[87OJs*dyN@w!]Q|f4Xu8Dko_)V~Xlw(x 6I7`mjM[JZ5-zY4M[s}F k' );

/**#@-*/

/**
 * WordPress database table prefix.
 *
 * You can have multiple installations in one database if you give each
 * a unique prefix. Only numbers, letters, and underscores please!
 */
$table_prefix = 'wp_';

/**
 * For developers: WordPress debugging mode.
 *
 * Change this to true to enable the display of notices during development.
 * It is strongly recommended that plugin and theme developers use WP_DEBUG
 * in their development environments.
 *
 * For information on other constants that can be used for debugging,
 * visit the documentation.
 *
 * @link https://wordpress.org/support/article/debugging-in-wordpress/
 */
define( 'WP_DEBUG', false );

/* Add any custom values between this line and the "stop editing" line. */



/* That's all, stop editing! Happy publishing. */

/** Absolute path to the WordPress directory. */
if ( ! defined( 'ABSPATH' ) ) {
	define( 'ABSPATH', __DIR__ . '/' );
}

/** Sets up WordPress vars and included files. */
require_once ABSPATH . 'wp-settings.php';
