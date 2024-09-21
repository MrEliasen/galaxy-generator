package consts

const (
	AU                        = 149.6e6                  // in km
	EARTH_RADII               = 6.371e6                  // in m
	EARTH_DENSITY             = 5.513                    // in g/m^3
	EARTH_MASS                = 5.97e24                  // in kg
	JOVIAN_MASS               = EARTH_MASS * 3.1782838e2 // in kg, used to estimate gas giant mass
	JOVIAN_RADII              = EARTH_RADII * 11.209     // in m, used to estimate gas giant mass
	NEPTUNE_MASS              = EARTH_MASS * 17.147      // in kg, used to estimate ice giant mass
	NEPTUNE_RADII             = EARTH_RADII * 3.883      // in m, used to estimate ice giant mass
	URANUS_MASS               = EARTH_MASS * 14.536      // in kg, used to estimate ice giant mass
	URANUS_RADII              = EARTH_RADII * 3.929      // in m, used to estimate ice giant mass
	SOLAR_MASS                = 1.989e30                 // in kg
	SOLAR_RADII               = 6.955e8                  // in meters
	SOLAR_WIND_MASS_LOSS_RATE = 1.4e12                   // kg/m^3
	SOLAR_WIND_VELOCITY       = 4.0e5                    // m/s
	SOLAR_ESCAPE_VELOCITY     = 617.54                   // km/s
	SOLAR_LUMINOSITY          = 3.828e26                 // solar lumens
	SOL_TEMP_K                = 5778.0                   // in kelvin
	GRAVITATIONAL_CONST       = 6.674e-11                // :)
	ISM_DENSITY               = 1.67e-21                 // kg/m^3
	ISM_VELOCITY              = 2.5e4                    // m/s
	STEFAN_BOLTZMANN_CONST    = 5.670374419e-8           // W/(m^2 K^4)
	BOLTZMANN_CONST           = 1.380649e-23             // J/K
	PROTON_MASS               = 1.6726219e-27            // in kg
	INNER_EDGE_STELLAR_FLUX   = 1.1                      // for a sun-like star
	OUTER_EDGE_STELLAR_FLUX   = 0.53                     // for a sun-like star
	INFRRED_FACTOR            = 0.2                      // Infrared adjustment factor for low-luminosity stars
	UV_FACTOR                 = 0.15                     // UV Penalty scaling factor for high-luminosity stars
	SOLAR_ABSOLUTE_MAGNITUDE  = 4.83                     // Absolute magnitude of the Sun
)
