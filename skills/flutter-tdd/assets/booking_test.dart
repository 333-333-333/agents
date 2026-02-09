// test/features/booking/domain/entities/booking_test.dart
import 'package:flutter_test/flutter_test.dart';
import 'package:bastet/features/booking/domain/entities/booking.dart';

void main() {
  group('Booking', () {
    group('creation', () {
      test('creates valid booking with all required fields', () {
        final booking = Booking(
          id: 'booking-1',
          ownerId: 'owner-1',
          caregiverId: 'caregiver-1',
          serviceType: ServiceType.walk,
          startAt: DateTime.now().add(const Duration(hours: 24)),
          endAt: DateTime.now().add(const Duration(hours: 25)),
          petIds: ['pet-1'],
        );

        expect(booking.id, equals('booking-1'));
        expect(booking.status, equals(BookingStatus.pending));
      });

      test('throws when no pets provided', () {
        expect(
          () => Booking(
            id: 'booking-1',
            ownerId: 'owner-1',
            caregiverId: 'caregiver-1',
            serviceType: ServiceType.walk,
            startAt: DateTime.now().add(const Duration(hours: 24)),
            endAt: DateTime.now().add(const Duration(hours: 25)),
            petIds: [],
          ),
          throwsArgumentError,
        );
      });

      test('throws when end is before start', () {
        expect(
          () => Booking(
            id: 'booking-1',
            ownerId: 'owner-1',
            caregiverId: 'caregiver-1',
            serviceType: ServiceType.walk,
            startAt: DateTime.now().add(const Duration(hours: 25)),
            endAt: DateTime.now().add(const Duration(hours: 24)),
            petIds: ['pet-1'],
          ),
          throwsArgumentError,
        );
      });
    });
  });
}
